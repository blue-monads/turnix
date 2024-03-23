<script>
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Home" />
</svelte:head>



<div class="flex justify-center p-10">
	<div class="relative overflow-x-auto">
		<table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
			<thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
				<tr>
					<th scope="col" class="px-6 py-3">
						Name
					</th>
					<th scope="col" class="px-6 py-3">
						Link
					</th>
				</tr>
			</thead>
			<tbody>
				<tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
					<th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
						Portal
					</th>
					<td class="px-6 py-4">
						<a href="/portal">click me</a>
					</td>
				</tr>
			   
				<tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
					<th scope="row" class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white">
						Login
					</th>
					<td class="px-6 py-4">
						<a href="/login">click me</a>
					</td>
				</tr>
			</tbody>
		</table>
	</div>
	
</div>



