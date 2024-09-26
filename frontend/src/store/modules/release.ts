import { head, uniq } from "lodash-es";
import { defineStore } from "pinia";
import { computed, reactive, ref, unref, watch } from "vue";
import { releaseServiceClient } from "@/grpcweb";
import type { MaybeRef, ComposedRelease, Pagination } from "@/types";
import { isValidDatabaseName, unknownRelease } from "@/types";
import { DEFAULT_PROJECT_NAME } from "@/types";
import type { Release } from "@/types/proto/v1/release_service";
import { DEFAULT_PAGE_SIZE } from "./common";
import { useProjectV1Store } from "./v1";
import { getProjectNameReleaseId, projectNamePrefix } from "./v1/common";

export const useReleaseStore = defineStore("release", () => {
  const releaseMapByName = reactive(new Map<string, ComposedRelease>());

  const releaseList = computed(() => {
    return Array.from(releaseMapByName.values());
  });

  const fetchReleasesByProject = async (
    project: string,
    pagination?: Pagination
  ) => {
    const resp = await releaseServiceClient.listReleases({
      parent: project,
      pageSize: pagination?.pageSize || DEFAULT_PAGE_SIZE,
      pageToken: pagination?.pageToken,
    });
    const composedReleaseList = await batchComposeRelease(resp.releases);
    composedReleaseList.forEach((release) => {
      releaseMapByName.set(release.name, release);
    });
    return {
      releases: composedReleaseList,
      nextPageToken: resp.nextPageToken,
    };
  };

  const fetchReleaseByName = async (name: string, silent = false) => {
    const release = await releaseServiceClient.getRelease({ name }, { silent });
    const [composedRelease] = await batchComposeRelease([release]);
    releaseMapByName.set(composedRelease.name, composedRelease);
    return composedRelease;
  };

  const getReleasesByProject = (project: string) => {
    return releaseList.value.filter((release) => release.project === project);
  };

  const getReleaseByName = (name: string) => {
    return releaseMapByName.get(name) ?? unknownRelease();
  };

  return {
    releaseList,
    fetchReleasesByProject,
    fetchReleaseByName,
    getReleasesByProject,
    getReleaseByName,
  };
});

export const useReleaseByName = (name: MaybeRef<string>) => {
  const store = useReleaseStore();
  const ready = ref(true);
  watch(
    () => unref(name),
    (name) => {
      if (!isValidDatabaseName(store.getReleaseByName(name).name)) {
        ready.value = false;
        store.fetchReleaseByName(name).then(() => {
          ready.value = true;
        });
      }
    },
    { immediate: true }
  );
  const release = computed(() => store.getReleaseByName(unref(name)));

  return {
    release,
    ready,
  };
};

export const batchComposeRelease = async (releaseList: Release[]) => {
  const composedReleaseList = releaseList.map((release) => {
    const composed = release as ComposedRelease;
    composed.project = `${projectNamePrefix}${head(getProjectNameReleaseId(release.name))}`;
    return composed;
  });
  const distinctProjectList = uniq(
    composedReleaseList.map((release) => release.project)
  );

  const projectV1Store = useProjectV1Store();
  await Promise.all(
    distinctProjectList.map((project) => {
      if (project === DEFAULT_PROJECT_NAME) {
        return;
      }
      return projectV1Store.getOrFetchProjectByName(project);
    })
  );
  return composedReleaseList.map((release) => {
    release.projectEntity = projectV1Store.getProjectByName(release.project);
    return release;
  });
};
