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

    <v-snackbar
      v-model="clipboardSnackbarShow"
      :color="clipboardSnackbarColor"
      :timeout="1500"
      bottom
    >
      {{ clipboardSnackbarText }}
    </v-snackbar>
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

  data: () => ({
    clipboardSnackbarShow: false,
  }),

  computed: mapState({
    matches: state => state.liveMatches.all,
    clipboardNotificationShow: state => state.liveMatches.clipboardNotification.show,
    clipboardSnackbarColor: state => {
      const notificationType = state.liveMatches.clipboardNotification.type;

      switch (notificationType) {
        case "success":
          return "primary";
        case "error":
          return "error";
        default:
          return "";
      }
    },
    clipboardSnackbarText: state => state.liveMatches.clipboardNotification.text,
  }),

  watch: {
    clipboardNotificationShow(val) {
      this.clipboardSnackbarShow = val;
    },
    clipboardSnackbarShow(val) {
      if (!val) {
        this.$store.commit("liveMatches/hideClipboardNotification");
      }
    },
  },
};
</script>
