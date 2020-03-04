<template>
  <v-container>
    <v-row align="stretch">
      <v-col
        v-for="match in matches"
        :key="match.match_id.toString()"
        cols="12"
        sm="6"
        md="4"
        lg="4"
        xl="3"
      >
        <LiveMatch :match="match" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import { mapState } from "vuex";
import LiveMatch from "@/components/LiveMatch.vue";

export default {
  name: "LiveMatches",

  components: {
    LiveMatch,
  },

  computed: mapState({
    matches: state => state.liveMatches.all,
  }),
};
</script>
