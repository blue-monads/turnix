<script>
	import {
		AppBar,
		AppRail,
		AppRailAnchor,
		AppRailTile,
		AppShell,
	} from "@skeletonlabs/skeleton";
	import "./styles.css";
	import Setting from "../compo/icons/Setting.svelte";

	let currentTile = "";

	const sibarItems = [
		{
			name: "Projects",
			icon: Setting,
		},
	];
</script>

<AppShell>
	<div slot="sidebarLeft" class="h-screen">
		<AppRail>
			<svelte:fragment slot="lead">
				<AppRailAnchor href="/">(icon)</AppRailAnchor>
			</svelte:fragment>

			{#each sibarItems as item, index}
				<AppRailTile
					bind:group={currentTile}
					name={item.name}

					value={index}
				>
					<svelte:fragment slot="lead">
						<svelte:component this={item.icon}></svelte:component>
					</svelte:fragment>

					<span>
						{item.name}
					</span>
				</AppRailTile>
			{/each}
		</AppRail>
	</div>

	<slot />
</AppShell>