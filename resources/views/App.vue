<template>
  <div>
    <vue-progress-bar/>
    <!-- Content -->
    <main-nav/>
    <div class="main-container">
      <main>
        <router-view/>
      </main>
    </div>
  </div>
</template>

<script>
  import MainNav from "@views/partials/MainNav"
  export default {
    components: {MainNav},
  };
</script>