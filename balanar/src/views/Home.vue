<template>
  <div class="home">
    <LiveMatches />
  </div>
</template>

<script>
import LiveMatches from "@/components/LiveMatches.vue";

export default {
  name: "Home",

  components: {
    LiveMatches,
  },
};
</script>
