<script lang="ts">
	import { onMount } from 'svelte';
	import { authenticated } from '../../../stores/auth';
	import { PUBLIC_SERVER_URL } from '$env/static/public';
	import type { Account } from '../../../../interfaces';

	let accountData: Account;

	onMount(async () => {
		const response = await fetch(`${PUBLIC_SERVER_URL}/api/accounts`, {
			headers: { 'Content-Type': 'application/json' },
			credentials: 'include'
		});

		const resOkay = response.ok;
		authenticated.set(resOkay);
		if (response.ok) {
			const data = await response.json();
			accountData = data;
		}
	});
</script>

{#if accountData}
	<div class="w-fit px-5 mx-auto my-10">
        <div class="text-xl my-20">
            <div class="text-center my-3">
                <label class="font-bold text-4xl" for="profile">Profile</label>
            </div>
            <div class="mx-auto my-3 w-64 h-64 rounded-full overflow-hidden">
                <!-- svelte-ignore a11y-img-redundant-alt -->
                <img src={accountData.image_uri} alt="profile-image" />
            </div>
            <div class="flex justify-between">
                <label class="font-bold w-1/2" for="username">Username</label>
                <div class="w-1/2">
                    {accountData.username}
                </div>
            </div>
            <div class="flex justify-between">
                <label class="font-bold w-1/2" for="email">Email</label>
                <div class="w-1/2">
                    {accountData.email}
                </div>
            </div>
            <div class="flex justify-between">
                <label class="font-bold w-1/2" for="gender">Gender</label>
                <div class="w-1/2">
                    {accountData.gender}
                </div>
            </div>
            <div class="flex justify-between">
                <label class="font-bold w-1/2" for="age">Age</label>
                <div class="w-1/2">
                    {accountData.age}
                </div>
            </div>
           
        </div>
    </div>
{/if}
