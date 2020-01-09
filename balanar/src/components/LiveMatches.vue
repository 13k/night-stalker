<template>
  <v-item-group>
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
          <v-item
            v-slot="{ active, toggle }"
            :value="match.match_id"
          >
            <LiveMatch
              :match="match"
              :active="active"
              :toggle="toggle"
            />
          </v-item>
        </v-col>
      </v-row>
    </v-container>
  </v-item-group>
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
