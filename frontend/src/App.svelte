<script lang="ts">
	export let name: string;
</script>

<svelte:head>
	<title>Fullstack App {name}</title>
</svelte:head>

<main class="container mx-auto max-w-screen-sm w-full">
	<h1
		class="space-y-8 mt-6 text-center text-3xl font-extrabold text-gray-900">
		Fullstack App
		{name}
	</h1>

	<div class="text-xs my-20">
		<em>Note: This is a template app for creating frontend apps quickly.
			This should not be a part of Production build.</em>
		<br /><br />
		The app is created with
		<a href="https://tailwindcss.com">Tailwind CSS</a>
		and
		<a href="https://svelte.dev">Svelte</a>.
	</div>
</main>
