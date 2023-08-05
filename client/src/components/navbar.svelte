<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { authenticated } from '../stores/auth';
	import { PUBLIC_SERVER_URL } from '$env/static/public';

	let image: string;
	let auth: boolean;
	authenticated.subscribe((isAuth) => (auth = isAuth));

	onMount(async () => {
		const response = await fetch(`${PUBLIC_SERVER_URL}/api/accounts`, {
			headers: { 'Content-Type': 'application/json' },
			credentials: 'include'
		});

		if (response.ok) {
			const data = await response.json();
			console.log(data)
			image = data.image_uri
		}
	});

	const logout = async () => {
		await fetch('http://localhost:3000/api/logout', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			credentials: 'include'
		});

    await goto("/login")
	};
</script>

<div class="navbar bg-base-100 shadow-sm">
	<div class="navbar-start">
		<div class="dropdown">
			<label tabIndex="0" for="hamburger" class="btn btn-ghost btn-circle">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					class="h-5 w-5"
					fill="none"
					viewBox="0 0 24 24"
					stroke="currentColor"
					><path
						stroke-linecap="round"
						stroke-linejoin="round"
						stroke-width="2"
						d="M4 6h16M4 12h16M4 18h7"
					/></svg
				>
			</label>
			<ul
				tabIndex="0"
				class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
			>
				<li><a href="/login">Login</a></li>
				<li><a href="/register">Register</a></li>
			</ul>
		</div>
	</div>
	<div class="navbar-center">
		<a href="/" class="btn btn-ghost normal-case text-xl">petplz</a>
	</div>
	<div class="navbar-end gap-1">
		<button class="btn btn-ghost btn-circle">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				class="h-5 w-5"
				fill="none"
				viewBox="0 0 24 24"
				stroke="currentColor"
				><path
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
				/></svg
			>
		</button>

		{#if auth}
			<div class="dropdown dropdown-end">
				<label tabIndex="0" for="account" class="btn btn-ghost btn-circle avatar">
					<div class="w-10 rounded-full">
						<img alt="user-icon"
							src={image}
						/>
					</div>
				</label>
				<ul
					tabIndex="0"
					class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
				>
					<li>
						<a href="/accounts/profile" class="justify-between">
							Profile
							<span class="badge">New</span>
						</a>
					</li>
					<li><button on:click={logout}>Logout</button></li>
				</ul>
			</div>
		{:else}
			<a class="btn btn-ghost" href="/login">login</a>
		{/if}
	</div>
</div>
