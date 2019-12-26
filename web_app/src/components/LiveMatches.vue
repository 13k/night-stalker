<template>
  <div class="live-matches">
    <live-match v-for="match in matches" :key="match.match_id" :match="match" />
  </div>
</template>

<script>
import { mapState } from "vuex";
import LiveMatch from "@/components/LiveMatch.vue";

export default {
  name: "live-matches",
  components: {
    LiveMatch
  },
  computed: mapState({
    matches: state => state.liveMatches.all
  })
};
</script>

<style lang="scss" scoped>
.live-matches {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
}
</style>
