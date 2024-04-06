<script lang="ts">
	import { API } from "$lib/api";
	import { getContext } from "svelte";
	import {
		AppBar,
		getModalStore,
		type ModalSettings,
	} from "@skeletonlabs/skeleton";
	import { FloatyButton } from "$lib/compo";
	import Icon from "@krowten/svelte-heroicons/Icon.svelte";

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
					>
					<Icon name="adjustments-horizontal" class="w-6 h-6" />
				</button>
			</footer>
		</div>
	{/each}
</div>

<FloatyButton
	handler={() => {
		store.trigger(modal);
	}}
/>
