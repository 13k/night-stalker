<template>
  <div class="home">
    <live-matches />
  </div>
</template>

<script>
import LiveMatches from "@/components/LiveMatches.vue";

export default {
  name: "home",
  components: {
    LiveMatches
  }
};
</script>
