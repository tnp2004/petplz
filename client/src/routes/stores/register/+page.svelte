<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_RECAPTCHA_SITE_KEY, PUBLIC_SERVER_URL } from '$env/static/public';

	let email: string;
	let password: string;

	let emailErrMsg: string;
	let passwordErrMsg: string;
	let errMsg: string;

	const submit = async () => {
		emailErrMsg = email ? '' : 'please enter your email';
		passwordErrMsg = password ? '' : 'please enter your password';

		if (!emailErrMsg && !passwordErrMsg) {
			const response = await fetch(`${PUBLIC_SERVER_URL}/api/login`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				credentials: 'include',
				body: JSON.stringify({
					email,
					password
				})
			});

			if (!response.ok) {
				errMsg = 'Email or Password is invalid';
				return;
			}

			// valid email and password
			errMsg = '';
			await goto('/');
		}
	};
</script>

<svelte:head>
	<script src="https://www.google.com/recaptcha/api.js" async defer></script>
</svelte:head>

<div class="flex">
	<form
		on:submit|preventDefault={submit}
		class="m-auto drop-shadow-xl bg-white rounded py-2 px-5 text-slate-700"
	>
		<h1 class="text-2xl font-bold bg-slate-200 w-fit px-3 py-1 rounded flex gap-1">
			<svg
				class="w-6"
				version="1.1"
				id="Layer_1"
				xmlns="http://www.w3.org/2000/svg"
				x="0"
				y="0"
				viewBox="0 0 32 32"
				style="enable-background:new 0 0 32 32"
				xml:space="preserve"
				><style>
					.st1 {
						fill: #333;
					}
				</style><path
					class="st1"
					d="M25.838 31H6.162a3.957 3.957 0 0 1-3.245-1.661 3.956 3.956 0 0 1-.549-3.604l.704-2.113a6.034 6.034 0 0 1 4.966-4.059C10.131 19.307 13.211 19 16 19c2.788 0 5.869.307 7.963.563a6.032 6.032 0 0 1 4.965 4.059l.704 2.113a3.954 3.954 0 0 1-.55 3.604A3.955 3.955 0 0 1 25.838 31zM16 21c-2.688 0-5.681.298-7.718.549a4.02 4.02 0 0 0-3.312 2.706l-.704 2.112c-.206.618-.106 1.274.274 1.802S5.511 29 6.162 29h19.676a1.98 1.98 0 0 0 1.622-.83c.381-.528.48-1.185.275-1.803l-.704-2.112a4.02 4.02 0 0 0-3.312-2.706C21.681 21.298 18.687 21 16 21zM16 18c-4.687 0-8.5-3.813-8.5-8.5S11.313 1 16 1c4.687 0 8.5 3.813 8.5 8.5S20.687 18 16 18zm0-15c-3.584 0-6.5 2.916-6.5 6.5S12.416 16 16 16s6.5-2.916 6.5-6.5S19.584 3 16 3z"
				/><path
					d="M12.04 10.54c-.543 0-.988-.435-1-.98a4.964 4.964 0 0 1 1.394-3.564 4.968 4.968 0 0 1 3.505-1.535c.562.01 1.009.428 1.02.98a1 1 0 0 1-.98 1.02 2.982 2.982 0 0 0-2.103.92 2.981 2.981 0 0 0-.836 2.139 1 1 0 0 1-.98 1.02h-.02z"
					style="fill:#008ad0"
				/></svg
			>
			ลงทะเบียนให้เช่า
		</h1>
		<div class="my-5 w-[304px]">
			<span hidden={!errMsg} class="text-red-600">*{errMsg}</span>
			<div class="my-2">
				<label for="email">
					ชื่อ
					<span hidden={!emailErrMsg} class="text-red-600">*{emailErrMsg}</span>
					<input bind:value={email} class="border-2 rounded-sm w-full px-1 my-2 h-9" type="text" />
				</label>

				<label for="description">
					คำอธิบาย
					<span hidden={!emailErrMsg} class="text-red-600">*{emailErrMsg}</span>
					<div>
						<textarea class="textarea textarea-bordered w-full" />
					</div>
				</label>

				<label for="rent-cost">
					ค่าเช่า
					<span hidden={!emailErrMsg} class="text-red-600">*{emailErrMsg}</span>
					<input bind:value={email} class="border-2 rounded-sm w-full px-1 my-2 h-9" type="text" />
				</label>

				<label for="type">
					ประเภท
					<span hidden={!passwordErrMsg} class="text-red-600">*{passwordErrMsg}</span>
					<input
						bind:value={password}
						class="border-2 rounded-sm w-full px-1 my-2 h-9"
						type="password"
					/>
				</label>

				<label for="rent-cost">
					รูปภาพ
					<span hidden={!emailErrMsg} class="text-red-600">*{emailErrMsg}</span>
					<input bind:value={email} class="rounded-sm w-full px-1 my-2 h-9" type="file" />
				</label>
			</div>
			<div class="g-recaptcha my-2 h-20" data-sitekey={PUBLIC_RECAPTCHA_SITE_KEY} />
			<button class="font-bold w-full btn btn-success rounded">เข้าสู่ระบบ</button>
		</div>
		<div class="text-center">
			<a class="text-slate-700/50" href="/register">อะไรนะ!! ไม่มีบัญชีหรอ ?</a>
		</div>
	</form>
</div>
