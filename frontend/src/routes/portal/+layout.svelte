<script>
	import {
		AppBar,
		AppRail,
		AppRailAnchor,
		AppRailTile,
		AppShell,
	} from "@skeletonlabs/skeleton";
	import Setting from "../../compo/icons/Setting.svelte";

	import Logo from  "../../compo/logo.png"

	let currentTile = "";

	const sibarItems = [
		{
			link: "/portal/projects",
			name: "Projects",
			icon: Setting,
		},
	];
</script>

<AppShell>
	<div slot="sidebarLeft" class="h-screen">
		<AppRail width="w-16">
			<svelte:fragment slot="lead">
				<AppRailAnchor href="/">

					<div class="flex justify-center items-center p-2">
						<img src={Logo} alt="logo">						
					</div>


				</AppRailAnchor>
			</svelte:fragment>

			{#each sibarItems as item, index}
				<AppRailTile
					bind:group={currentTile}
					name={item.name}

					value={index}
				>
					<svelte:fragment slot="lead">
						<a class="flex flex-col justify-center items-center" href={item.link}>
							<svelte:component this={item.icon}></svelte:component>
							<span class="text-xs">
								{item.name}
							</span>
						</a>
					</svelte:fragment>


				</AppRailTile>
			{/each}
		</AppRail>
	</div>

	<slot />
</AppShell>