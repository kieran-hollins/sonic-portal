<script lang="ts">
    const title: string[] = "PORTALS".split("");
</script>

<div class="hero-title-container">
    <h1 class="hero-title">
        {#each title as char, i}
            <span class="letter">{char}</span>
        {/each}
    </h1>
</div>

<style>
    .hero-title {
        display: flex;
        justify-content: center;
        font-size: 6rem;
        font-family: "Space Mono", monospace;
        perspective: 1000px;
        overflow: hidden;
    }

    .letter {
        color: white;
        display: inline-block;
        opacity: 0;
        transform: rotateY(-90deg);
        
        animation-name: appear, fadeOut;
        animation-duration: 0.8s, 2s;
        animation-timing-function: ease-out, ease-in;
        animation-fill-mode: forwards, forwards;
        animation-delay: var(--appear-delay), var(--fadeout-delay);
    }

    .letter:nth-child(1) { animation-delay: 0s, 3s; }
    .letter:nth-child(2) { animation-delay: 0.1s, 3.1s; }
    .letter:nth-child(3) { animation-delay: 0.2s, 3.2s; }
    .letter:nth-child(4) { animation-delay: 0.3s, 3.3s; }
    .letter:nth-child(5) { animation-delay: 0.4s, 3.4s; }
    .letter:nth-child(6) { animation-delay: 0.5s, 3.5s; }
    .letter:nth-child(7) { animation-delay: 0.6s, 3.6s; }

    @keyframes appear {
        from {
            opacity: 0;
            transform: rotateY(-90deg);
        }
        to {
            opacity: 1;
            transform: rotateY(0deg);
        }
    }

    @keyframes fadeOut {
        from {
            opacity: 1;
        }
        to {
            opacity: 0;
        }
    }
</style>