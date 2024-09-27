import { json } from '@sveltejs/kit';

/** @type {import('./$types').RequestHandler} */
export async function GET({ route }) {
	console.log(`BACKEND_URL: ${import.meta.env.VITE_BACKEND_URL}`)
	console.log(`url: ${route.id}`)
	const request = await fetch(`${import.meta.env.VITE_BACKEND_URL}${route.id}`)
	const result = await request.json();
	return json(result);
}