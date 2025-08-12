<!-- -->

<script>
    import Header from "$lib/components/Header.svelte";
    import HeroTitle from "$lib/components/HeroTitle.svelte";
    import Portals from "$lib/components/Portals.svelte";
    import '../app.css';
</script>

<div class="container">
    <Header />
    <HeroTitle />
    <Portals />
</div>

<style>    
  .container {
    background-color: #030303;
    width: 100vw;
    height: 100vh;

    /* display: flex; */
    /* align-content: center; */
    /* justify-content: center; */
  }
</style>