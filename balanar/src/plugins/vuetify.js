import Vue from "vue";
import Vuetify from "vuetify/lib";

import colors from "vuetify/lib/util/colors";

Vue.use(Vuetify);

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: colors.purple.darken1,
        secondary: colors.grey.darken1,
        accent: colors.shades.black,
        error: colors.red.accent3,
      },
      dark: {
        primary: colors.deepPurple.base,
        secondary: colors.purple.base,
        accent: colors.pink.base,
        error: colors.red.base,
        warning: colors.amber.base,
        info: colors.cyan.base,
        success: colors.lightGreen.base,
      },
    },
  },
});
