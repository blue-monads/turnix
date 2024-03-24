<script lang="ts">
	import {
		AppBar,
		AppRail,
		AppRailAnchor,
		AppRailTile,
		AppShell,
	} from "@skeletonlabs/skeleton";
	import Logo from "../../lib/images/logo.png";
	import { onMount } from "svelte";
	import Icon from "@krowten/svelte-heroicons/Icon.svelte";
	import { API } from "$lib/api/api";
	import ContextThis from "./contextThis.svelte";

	let currentTile = "";

	const sibarItems = [
		{
			link: "/portal/projects",
			name: "Projects",
			icon: "rectangle-group",
		},

		{
			link: "/portal/setting",
			name: "Setting",
			icon: "adjustments-horizontal",
		},
	];

	let api: API;

	const load = () => {
		api = new API()

	}

	load()

</script>

<AppShell>
	<div slot="sidebarLeft" class="h-screen">
		<AppRail width="w-14">
			<div slot="lead" class="mb-4">
				<AppRailAnchor href="/">
					<div class="flex justify-center items-center p-2">
						<img src={Logo} alt="logo" />
					</div>
				</AppRailAnchor>
			</div>

			{#each sibarItems as item, index}
				<AppRailTile
					bind:group={currentTile}
					name={item.name}
					value={index}
				>
					<svelte:fragment slot="lead">
						<a
							class="flex flex-col justify-center items-center"
							href={item.link}
						>
							<Icon name={item.icon} class="w-6 h-6" />
							<span class="text-xs">
								{item.name}
							</span>
						</a>
					</svelte:fragment>
				</AppRailTile>
			{/each}
		</AppRail>
	</div>

	<ContextThis {api}>
		<svelte:fragment>
			<slot />
		</svelte:fragment>
	</ContextThis>
</AppShell>
