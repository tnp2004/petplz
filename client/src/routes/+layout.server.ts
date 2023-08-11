export function load({ cookies }) {
	const jwt = cookies.get('jwt');

	return {
		jwt
	};
}
