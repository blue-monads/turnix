<script lang="ts">
	// @ts-ignore
	import Icon from "@krowten/svelte-heroicons/Icon.svelte";
	import {
		AppBar,
		AppRail,
		AppRailAnchor,
		AppRailTile,
		AppShell,
		type PopupSettings,
	} from "@skeletonlabs/skeleton";
	import { popup } from "@skeletonlabs/skeleton";
	import {
		computePosition,
		autoUpdate,
		offset,
		shift,
		flip,
		arrow,
	} from "@floating-ui/dom";

	import { storePopup } from "@skeletonlabs/skeleton";

	import Logo from "../../lib/images/logo.png";
	import { RootAPI } from "$lib/api/api";
	import ContextThis from "./contextThis.svelte";

	let currentTile = "";
	let showMobileModal = false;

	let sibarItems: Record<string, any>[] = [];

	let api: RootAPI;

	const load = () => {
		// @ts-ignore
		const ptypes: object[] = window["__turnis_ptypes__"] || [];
		sibarItems = ptypes;

		api = new RootAPI();
	};

	const toggle = () => {
		showMobileModal = !showMobileModal;
	};

	const popupHover: PopupSettings = {
		event: "hover",
		target: "popupHover",
		placement: "right",
	};

	storePopup.set({ computePosition, autoUpdate, offset, shift, flip, arrow });

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

					<li class="border rounded">
						<a
							on:click={toggle}
							class="h-12 px-4 flex justify-center items-center w-full text-gray-700 focus:text-orange-500 uppercase gap-2"
							href="/z/pages/portal">
							<Icon
								name="home"
								class="w-6 h-6"
							/>
							Home
						</a>
					</li>


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
	<div
		slot="header"
		class="flex flex-row md:hidden justify-between bg-surface-100-800-token"
	>
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

	<div
		slot="sidebarLeft"
		class="h-screen hidden md:block bg-gradient-to-br variant-gradient-primary-secondary"
	>
		<AppRail width="w-14">
			<div slot="lead" class="mb-4">
				<AppRailAnchor href="/z/pages/portal">
					<div class="flex justify-center items-center p-2">
						<img src={Logo} alt="logo" />
					</div>
				</AppRailAnchor>
			</div>

			<div class="flex flex-col">
				{#each sibarItems as item, index}
				{@const popupHover = { event: "hover", target: `popupHover-${index}`,placement: "right",}}
					<a
					
						use:popup={popupHover}
						class="p-4 hover:bg-secondary-50 [&>*]:pointer-events-none"
						href={item["link"]}
					>
						<Icon name={item["icon"] || "link"} class="w-6 h-6" />
					</a>

					<div class="card p-2 variant-filled-secondary" data-popup={`popupHover-${index}`}>
						<p>{item["name"] || ""}</p>
					</div>

				{/each}
			</div>

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
