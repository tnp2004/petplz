<script lang="ts">
	import { onMount } from 'svelte';
	import { authenticated } from '../../../stores/auth';

	interface Account {
		accountId: string;
		username: string;
		email: string;
		gender: 'male' | 'female' | 'other';
		age: number;
		money: number;
		image_uri: string;
	}

	let accountData: Account;

	onMount(async () => {
		const response = await fetch('http://localhost:3000/api/accounts', {
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
        <div class="text-xl my-20 w-56">
            <div class="text-center my-3">
                <label class="font-bold text-4xl" for="profile">Profile</label>
            </div>
            <div class="my-3">
                <img class="rounded-full" src={accountData.image_uri} alt="profile-image" />
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
