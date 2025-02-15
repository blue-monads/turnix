<script lang="ts">
    import { goto } from "$app/navigation";
    import { AutoTable, FloatyButton } from "$lib/compo";
    import SvgIcon from "$lib/compo/icons/SvgIcon.svelte";
    import { getContext } from "svelte";
    import { AppBar, getModalStore } from "@skeletonlabs/skeleton";
    import type { RootAPI } from "$lib/api";
    import { page } from "$app/stores";
    import { NewZblogAPI } from "../../zblogApi";
  
    const pid = $page.params["pid"];
  
    const store = getModalStore();
    const api = NewZblogAPI(getContext("__api__") as RootAPI);
  
    let posts: any[] = $state([]);
  
    const load = async () => {
      const resp = await api.listPost(pid);
      if (resp.status !== 200) {
        return;
      }
  
      posts = resp.data;
    };
  
    load();
  </script>
  
  <AppBar>
    <svelte:fragment slot="lead">
      <ol class="breadcrumb">
        <li class="crumb">
          <a class="anchor" href="/z/pages/portal/project/list/zblog">ZBlogs</a>
        </li>
        <li class="crumb-separator" aria-hidden>&rsaquo;</li>
        <li>Posts</li>
      </ol>
    </svelte:fragment>
  
    <div class="flex flex-wrap justify-end gap-2">
      <a
        href={`/z/pages/portal/projects/zblog/${pid}/posts/new`}
        class="btn variant-filled btn-sm"
      >
        <SvgIcon className="h-4 w-4" name="plus" />
  
        Post
      </a>
  
      <a
        href={`/z/pages/portal/projects/zblog/${pid}/tags/new`}
        class="btn variant-filled-secondary btn-sm"
      >
        <SvgIcon className="h-4 w-4" name="plus" />
  
        Tag
      </a>
  
  
    </div>
  </AppBar>
  
  <AutoTable
    action_key={"id"}
    key_names={[
      ["id", "ID"],
      ["title", "Title"],
      ["slug", "Slug"],
      ["excerpt", "Excerpt"],
      ["created_at", "Created At"],

    ]}
    datas={posts}

    actions={[  
      {
        Name: "edit",
        Class: "variant-filled-secondary",
        Action: async (id) => {

        },
      },
    ]}
  />
  