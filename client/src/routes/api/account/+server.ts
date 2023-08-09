import { json } from '@sveltejs/kit';

export async function GET() {
    try {
        const response = await fetch("http://localhost:3000/api/accounts")
        const data = await response.json()
        return json(data);
    } catch (e) {
        return e
    }

}
