<script lang="ts">
	import { API } from "$lib/api";
	import { getContext } from "svelte";
	import { AppBar, getModalStore, type ModalSettings } from "@skeletonlabs/skeleton";

	const api = getContext("__api__") as API;

	let projects = [
		{
			id: 12,
			ptype: "xyz",
			name: "test",
			info: "This is test",
			isInitilized: false,
		},
	];

	const load = async () => {
		const resp = await api.listProjects();
		if (resp.status !== 200) {
			return;
		}

		projects = resp.data;
	};

	const modal: ModalSettings = {
		type: "component",
		component: "picker",
	};

	const store = getModalStore();

	load();
</script>

<svelte:head>
	<title>Projects</title>
</svelte:head>

<AppBar>
	<svelte:fragment slot="lead">
        <h4 class="h4">Home</h4>
    </svelte:fragment>

</AppBar>


<div class="p-4 flex flex-wrap gap-2">
	{#each projects as proj}
		<div class="card p-1 w-96 hover:bg-secondary-backdrop-token">
			<header class="card-header">
				<img
					src={`https://picsum.photos/seed/${proj.id}/512/300`}
					alt=""
				/>
			</header>
			<section class="p-4 flex flex-col">
				<div class="flex justify-end">
					<span
						class="chip variant-filled-tertiary bg-tertiary-hover-token"
						>{proj.ptype}</span
					>
				</div>

				<h3 class="h3">{proj.name}</h3>

				<p class="p">{proj.info}</p>
			</section>

			<footer class="card-footer flex justify-end gap-2">
				<button class="btn variant-filled">explore</button>
				<button class="btn variant-filled variant-filled-warning"
					>delete</button
				>
			</footer>
		</div>
	{/each}
</div>

<div class="fixed bottom-5 right-5">
	<button
		class="btn-icon variant-filled bg-primary-hover-token"
		on:click={() => {
			store.trigger(modal);
		}}
	>
		<svg viewBox="0 0 24 24" fill="currentColor" class="w-6 h-6">
			<path
				fill-rule="evenodd"
				d="M12 3.75a.75.75 0 0 1 .75.75v6.75h6.75a.75.75 0 0 1 0 1.5h-6.75v6.75a.75.75 0 0 1-1.5 0v-6.75H4.5a.75.75 0 0 1 0-1.5h6.75V4.5a.75.75 0 0 1 .75-.75Z"
				clip-rule="evenodd"
			/>
		</svg>
	</button>
</div>
