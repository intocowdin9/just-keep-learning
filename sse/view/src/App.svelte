<script>
	import { onMount } from "svelte"
	let time = ""
	let whoami = ""
	let displayTime = false;
	let displayUser = false;

	onMount( () => {
		const evtSrc = new EventSource("http://127.0.0.1:9000/event")
		
		evtSrc.onmessage = function(event) {

			time = event.data
		}

		evtSrc.onerror = function(event) {
			console.log(event)
		}

	})

	async function getTime() {
		const res = await fetch("http://localhost:9000/time")
		if (res.status !== 200) {
			console.log("could not connect to the server")
		} else {
			console.log("OK")
		}

		displayTime = true
		displayUser = false
	}

	async function getUser() {
		const res = await fetch("http://localhost:9000/user")
		if (res.status !== 200) {
			console.log("could not connect to the server")
		} else {
			console.log("OK")
		}

		displayTime = false
		displayUser = true
	}
</script>

<main>
	<h1>Server Sent Events</h1>
	<button on:click={getTime}>Get Time</button>
	<button on:click={getUser}>Get User</button>
	{#if displayTime}
    <p>{time}</p>
  {/if}
  {#if displayUser}
    <p>{whoami}</p>
  {/if}
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>