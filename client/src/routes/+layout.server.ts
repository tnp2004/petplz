import type { LayoutServerLoad } from "./$types";

export const load: LayoutServerLoad = ({ cookies }) => {
	const jwt = cookies.get('jwt');

	return {
		jwt
	};
}
