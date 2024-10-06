<script lang="ts">
	import { AppBar } from "@skeletonlabs/skeleton";
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Home" />
</svelte:head>

<div class="w-full h-full">
	<AppBar>
		<svelte:fragment slot="lead">Turnix</svelte:fragment>
	</AppBar>


	<div class="my-10 mx-auto w-96 p-4">
	
		<div class="flex flex-col gap-2 card p-4">

			<h3 class="h4">Pages</h3>

			<div class="flex flex-col pl-2">

				<ul class="list-disc list-inside">
					<li><a class="hover:underline" href="/z/pages/portal">Portal</a></li>
					<li><a class="hover:underline" href="/z/pages/auth/login">Login</a></li>
					<li><a class="hover:underline" href="/z/pages/auth/login">Signup</a></li>
				</ul>

			</div>
		</div>
	
	</div>

	



</div>
