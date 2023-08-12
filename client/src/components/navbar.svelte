<script lang="ts">
	import { goto, invalidate } from '$app/navigation';
	import { onMount } from 'svelte';
	import { authenticated } from '../stores/auth';
	import { PUBLIC_SERVER_URL } from '$env/static/public';
	import type { LayoutData } from '../routes/$types';

	let image: string;
	let auth: boolean;
	export let data: LayoutData;
	authenticated.subscribe((isAuth) => (auth = isAuth));

	onMount(async () => {
		if (data.jwt) {
			const response = await fetch(`${PUBLIC_SERVER_URL}/api/accounts`, {
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include'
			});

			if (response.ok) {
				const data = await response.json();
				image = data.image_uri;
			}
		}
	});

	const refresh = () => {
		invalidate(`${PUBLIC_SERVER_URL}/api/accounts`)
	}
	
	const logout = async () => {
		await fetch(`${PUBLIC_SERVER_URL}/api/logout`, {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			credentials: 'include'
		});
		await refresh()
		await goto('/login');
	};
</script>

<div class="navbar bg-base-100 shadow-sm mb-8 text-slate-700">
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
				<li><a href="/stores">ร้านค้า</a></li>
				<li><a href="/stores/register">ลงทะเบียน</a></li>
			</ul>
		</div>
	</div>
	<div class="navbar-center gap-2">
		<div class="menu menu-horizontal px-1 gap-1">
			<a href="/" class="btn btn-ghost normal-case text-2xl">petplz</a>
		</div>
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
						<img alt="user-icon" src={image} />
					</div>
				</label>
				<ul
					tabIndex="0"
					class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
				>
					<li>
						<a href="/accounts/profile" class="justify-between">
							โปรไฟล์
							<span class="badge">ใหม่</span>
						</a>
					</li>
					<li><button on:click={logout}>ออกจากระบบ</button></li>
				</ul>
			</div>
		{:else}
			<a class="btn btn-ghost" href="/login">เข้าสู่ระบบ</a>
		{/if}
	</div>
</div>
