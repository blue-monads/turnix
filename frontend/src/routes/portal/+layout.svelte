<script lang="ts">
	import Icon from "@krowten/svelte-heroicons/Icon.svelte";
	import {
		AppBar,
		AppRail,
		AppRailAnchor,
		AppRailTile,
		AppShell,
	} from "@skeletonlabs/skeleton";
	import Logo from "../../lib/images/logo.png";

	import { API } from "$lib/api/api";
	import ContextThis from "./contextThis.svelte";

	let currentTile = "";
	let showMobileModal = false;

	let sibarItems: object[] = [];

	let api: API;

	const load = () => {
		const ptypes: object[] = window["__turnis_ptypes__"] || [];
		sibarItems = ptypes;

		api = new API();
	};

	const toggle = () => {
		showMobileModal = !showMobileModal;
	};

	load();
</script>

{#if showMobileModal}
	<div class="h-screen w-screen bg-gray-600 bg-opacity-90 fixed z-50">
		<div
			class="h-full w-full absolute transform translate-x-0 transition ease-in-out duration-700 p-5"
		>
			<div class="h-full bg-white rounded p-5">
				<div class="absolute right-4">
					<button
						on:click={toggle}
						aria-label="close menu modal"
						class="focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 cursor-pointer rounded border mr-2"
					>
						<Icon name="x-mark" class="h-4 w-4" />
					</button>
				</div>

				<ul class="space-y-4 pt-10">
					{#each sibarItems as item, index}
						<li class="border rounded">
							<a
								on:click={toggle}
								class="h-12 px-4 flex justify-center items-center w-full text-gray-700 focus:text-orange-500 uppercase gap-2"
								href={item["link"]}
							>
								<Icon
									name={item["icon"] || "link"}
									class="w-6 h-6"
								/>
								{item["name"] || ""}
							</a>
						</li>
					{/each}
				</ul>
			</div>
		</div>
	</div>
{/if}

<AppShell>
	<div slot="header" class="flex flex-row md:hidden justify-between bg-surface-100-800-token">
		<div class="flex p-1">
			<button on:click={toggle}>
				<Icon name="bars-3" class="w-6 h-6" />
			</button>
		</div>

		<div class="flex gap-4 p-1">
			<a
				class="flex flex-col justify-center items-center rounded-full bg-secondary-100 hover:bg-secondary-200 p-2"
				href={"/z/pages/portal/self"}
			>
				<Icon name="user" class="w-6 h-6" />
			</a>
		</div>
	</div>

	<div slot="sidebarLeft" class="h-screen hidden md:block">
		<AppRail width="w-14">
			<div slot="lead" class="mb-4">
				<AppRailAnchor href="/z/pages/portal">
					<div class="flex justify-center items-center p-2">
						<img src={Logo} alt="logo" />
					</div>
				</AppRailAnchor>
			</div>

			{#each sibarItems as item, index}
				<AppRailTile
					bind:group={currentTile}
					name={item["name"]}
					value={index}
				>
					<svelte:fragment slot="lead">
						<a
							class="flex flex-col justify-center items-center"
							href={item["link"]}
						>
							<Icon
								name={item["icon"] || "link"}
								class="w-6 h-6"
							/>
						</a>
					</svelte:fragment>
				</AppRailTile>
			{/each}

			<div slot="trail" class="mb-4">
				<a
					class="flex flex-col justify-center items-center rounded-full bg-secondary-100 hover:bg-secondary-200 p-2"
					href={"/z/pages/portal/self"}
				>
					<Icon name="user" class="w-6 h-6" />
				</a>
			</div>
		</AppRail>
	</div>

	<ContextThis {api}>
		<svelte:fragment>
			<slot />
		</svelte:fragment>
	</ContextThis>
</AppShell>
