<script lang="ts">
	let currentPortal = 0;
	const portals = Array(5).fill(null).map(() => randomGradient());

	const displayTime = 5000; 
	const animTime = 1000; 

	function randomGradient() {
		const c1 = `hsl(${Math.random() * 360}, 80%, 60%)`;
		const c2 = `hsl(${Math.random() * 360}, 80%, 60%)`;
		return `radial-gradient(circle, ${c1} 0%, ${c2} 100%)`;
	}

	function nextPortal() {
		const oldPortal = currentPortal;
		portals[oldPortal] = portals[oldPortal];

		setTimeout(() => {
			currentPortal = (currentPortal + 1) % portals.length;
			portals[currentPortal] = randomGradient();
		}, animTime);
	}

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
    filter: blur(20px);
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
