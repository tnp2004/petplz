<svelte:head>
	<script src="https://www.google.com/recaptcha/api.js" async defer></script>
</svelte:head>

<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_RECAPTCHA_SITE_KEY, PUBLIC_SERVER_URL } from '$env/static/public';

	// Value
	let username: string;
	let email: string;
	let password: string;
	let confirmPassword: string;
	let age: number;
	let gender: string;

	// Error message
	let usernameErrMsg: string;
	let emailErrMsg: string;
	let passwordErrMsg: string;
	let confirmPasswordErrMsg: string;
	let ageErrMsg: string;
	let genderErrMsg: string;

	const usernameOkay = async (): Promise<boolean> => {

		if (!username) {
			usernameErrMsg = "กรุณากรอกชื่อผู้ใช้"
			return false
		}

		const response = await fetch(`${PUBLIC_SERVER_URL}/api/validate?username=${username}`)

		if (!response.ok) {
			usernameErrMsg = "ชื่อนี้ถูกใช้ไปแล้ว"
			return false;
		}

		usernameErrMsg = ""
		return true
	};

	const emailOkay = async (): Promise<boolean> => {

		if (!email) {
			emailErrMsg = "กรุณากรอกอีเมล"
			return false
		}

		const response = await fetch(`${PUBLIC_SERVER_URL}/api/validate?email=${email}`)

		if (!response.ok) {
			emailErrMsg = "อีเมลนี้ถูกใช้ไปแล้ว"
			return false;
		}

		emailErrMsg = ""
		return true
	};

	const passwordOkay = (): boolean => {

		if (!password) {
			passwordErrMsg = "กรุณากรอกรหัสผ่าน"
			return false
		}else if (password.length < 5) {
			passwordErrMsg = "รหัสผ่านต้องมีอย่างน้อย 5 ตัว"
			return false
		} else {
			passwordErrMsg = ""
		}
		
		if (password === confirmPassword) {
			confirmPasswordErrMsg = ""
			return true
		}

		confirmPasswordErrMsg = "รหัสผ่านไม่ตรงกัน"
		return false;
	};

	const ageOkay = (): boolean => {
		if (age) {
			ageErrMsg = ""
			return true
		}

		ageErrMsg = "กรุณากรอกอายุ"
		return false
	}

	const genderOkay = (): boolean => {
		if (gender) {
			genderErrMsg = ""
			return true
		}

		genderErrMsg = "กรุณาเลือกเพศ"
		return false
	}

	const submit = async () => {
		const canSubmit = await usernameOkay() && await emailOkay() && passwordOkay() && ageOkay() && genderOkay();
		if (canSubmit) {
			const response =  await fetch(`${PUBLIC_SERVER_URL}/api/accounts/register`, {
				method: 'POST',
				headers: { 'Content-Type': 'application/json' },
				body: JSON.stringify({
					username,
					email,
					password,
					gender,
					age
				})
			});

			if (response.ok) {
				await goto('/login');
			}
		}
	};
</script>

<div class="flex mb-5">
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
			สมัครสมาชิก
		</h1>
		<div class="my-5">
			<div class="my-2 w-[304px]">
				<label for="username">
					ชื่อผู้ใช้
					<span hidden={!usernameErrMsg} class="text-red-600">*{usernameErrMsg}</span>
					<input
						bind:value={username}
						class={`border-2 rounded-sm w-full px-1 my-2 h-9 ${usernameErrMsg ? "error-input": ""}`}
						type="text"
					/>
				</label>
				<label for="email">
					อีเมล
					<span hidden={!emailErrMsg} class="text-red-600">*{emailErrMsg}</span>
					<input
						bind:value={email}
						class={`border-2 rounded-sm w-full px-1 my-2 h-9 ${emailErrMsg ? "error-input": ""}`}
						type="email"
					/>
				</label>
				<label for="password">
					รหัสผ่าน
					<span hidden={!passwordErrMsg} class="text-red-600">*{passwordErrMsg}</span>
					<input
						bind:value={password}
						class={`border-2 rounded-sm w-full px-1 my-2 h-9 ${passwordErrMsg ? "error-input": ""}`}
						type="password"
					/>
				</label>
				<label for="confirm-password">
					ยืนยันรหัสผ่าน
					<span hidden={!confirmPasswordErrMsg} class="text-red-600">*{confirmPasswordErrMsg}</span>
					<input
						bind:value={confirmPassword}
						class={`border-2 rounded-sm w-full px-1 my-2 h-9 ${confirmPasswordErrMsg ? "error-input": ""}`}
						type="password"
					/>
				</label>
				<label for="age">
					อายุ
					<span hidden={!ageErrMsg} class="text-red-600">*{ageErrMsg}</span>
					<input bind:value={age} class={`border-2 rounded-sm w-full px-1 my-2 h-9 ${ageErrMsg ? "error-input": ""}`} type="number" />
				</label>

				<span hidden={!genderErrMsg} class="text-red-600">*{genderErrMsg}</span>
				<div class="flex gap-2 my-2 justify-between">
					<button
						on:click={() => (gender = 'male')}
						class="rounded w-full border-2 p-1 flex justify-around"
						type="button"
					>
						<input
							class="radio radio-success w-5 h-5 my-auto mx-1"
							type="radio"
							value="male"
							checked={gender == 'male'}
							required
						/>
						<span class="w-full">ชาย</span>
					</button>
					<button
						on:click={() => (gender = 'female')}
						class="rounded w-full border-2 p-1 flex justify-around"
						type="button"
					>
						<input
							class="radio radio-success w-5 h-5 my-auto mx-1"
							type="radio"
							value="female"
							checked={gender == 'female'}
						/>
						<span class="w-full">หญิง</span>
					</button>
					<button
						on:click={() => (gender = 'other')}
						class="rounded w-full border-2 p-1 flex justify-around"
						type="button"
					>
						<input
							class="radio radio-success w-5 h-5 my-auto mx-1"
							type="radio"
							value="other"
							checked={gender == 'other'}
						/>
						<span class="w-full">อื่นๆ</span>
					</button>
				</div>
			</div>
			<div class="g-recaptcha my-2 h-20" data-sitekey={PUBLIC_RECAPTCHA_SITE_KEY}></div>
			<button type="submit" class="font-bold w-full btn btn-success rounded">SIGN UP</button>
		</div>
		<div class="text-center text-slate-400">
			<a class="text-slate-700/50" href="/login">มีบัญชีแล้วงั้นหรอ ?</a>
		</div>
	</form>
</div>
