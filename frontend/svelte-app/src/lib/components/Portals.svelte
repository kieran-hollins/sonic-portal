<script lang="ts">
	let currentPortal = 0;
	const portals = Array(5).fill(null).map(() => randomGradient());

	// Animation timing
	const displayTime = 2000; // time portal stays visible before fading
	const animTime = 1000; // must match CSS animation durations

	function randomGradient() {
		// Generates a radial gradient with two random colors
		const c1 = `hsl(${Math.random() * 360}, 80%, 60%)`;
		const c2 = `hsl(${Math.random() * 360}, 80%, 60%)`;
		return `radial-gradient(circle, ${c1} 0%, ${c2} 100%)`;
	}

	function nextPortal() {
		// Trigger arcOut on the current portal
		const oldPortal = currentPortal;
		portals[oldPortal] = portals[oldPortal]; // no change, just here for clarity

		// Wait for arcOut to finish, then switch & randomize
		setTimeout(() => {
			currentPortal = (currentPortal + 1) % portals.length;
			portals[currentPortal] = randomGradient();
		}, animTime);
	}

	// Loop through portals
	setInterval(nextPortal, displayTime);
</script>

<div class="portal-container">
	{#each portals as bg, i}
		<div
			class="portal 
				{i === currentPortal ? 'arcIn' : 'arcOut'}"
			style="background: {bg};"
		></div>
	{/each}
</div>

<style>
.portal-container {
	width: 100%;
	height: 100%;
	max-width: 1200px;
	max-height: 1200px;
	aspect-ratio: 1 / 1;
	display: flex;
	position: relative;
	align-content: center;
	justify-content: center;
	overflow: hidden;
}

.portal {
	border-radius: 50%;
	height: 400px;
	width: 400px;
	position: absolute;
	opacity: 0;
	transform: translate3d(-480px, 480px, 0);
    filter: blur(3px);
}

@keyframes arcIn {
	from {
		opacity: 0;
		transform: translate3d(-480px, 480px, 0);
	}
	to {
		opacity: 1;
		transform: translate3d(0, 0, 0);
	}
}

@keyframes arcOut {
	from {
		opacity: 1;
		transform: translate3d(0, 0, 0);
	}
	to {
		opacity: 0;
		transform: translate3d(480px, 480px, 0);
	}
}

.arcIn {
	animation: arcIn 1s ease forwards;
}

.arcOut {
	animation: arcOut 1s ease forwards;
}
</style>
