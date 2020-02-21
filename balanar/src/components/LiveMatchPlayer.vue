<template>
  <div
    class="live-match-player d-flex align-center"
    :class="containerClasses"
  >
    <div
      class="d-flex"
      :class="nameClasses"
    >
      <span>{{ player.name }}</span>
      <span
        v-if="label"
        class="ml-2"
      >
        ({{ label }})
      </span>

      <span
        v-if="kda"
        class="ml-2 overline grey--text"
      >
        {{ kda.kills }}/{{ kda.deaths }}/{{ kda.assists }}
      </span>
    </div>

    <div
      class="d-flex"
      :class="iconClasses"
    >
      <HeroImage
        :hero="hero"
        version="portrait"
        width="64"
        height="36"
        placeholder
        alt
        alt-placeholder="Picking hero"
      />

      <div
        class="slot-bar"
        :class="slotBarClasses"
      />
    </div>
  </div>
</template>

<script>
import _ from "lodash";

import * as t from "@/protocol/transform";
import HeroImage from "@/components/HeroImage.vue";

const cleanName = name =>
  _.chain((name || "").normalize())
    .deburr()
    .toLower()
    .replace(/\W/g, "-")
    .replace(/-+/g, "-")
    .replace(/(^-|-$)/, "")
    .value();

export default {
  name: "LiveMatchPlayer",

  components: {
    HeroImage,
  },

  props: {
    team: {
      type: Object,
      required: true,
    },
    player: {
      type: Object,
      required: true,
    },
  },

  computed: {
    hero() {
      return t.get(this.player, "hero");
    },
    slot() {
      return t.get(this.player, "slot");
    },
    label() {
      const label = cleanName(this.player.label);
      const name = cleanName(this.player.name);

      if (name !== label) {
        return this.player.label;
      }

      return null;
    },
    kda() {
      if (this.player.kills > 0 || this.player.deaths > 0 || this.player.assists > 0) {
        return {
          kills: this.player.kills || 0,
          deaths: this.player.deaths || 0,
          assists: this.player.assists || 0,
        };
      }

      return null;
    },
    isLeft() {
      return this.team.number % 2 === 0;
    },
    isRight() {
      return !this.isLeft;
    },
    containerClasses() {
      return {
        "justify-end": this.isRight,
      };
    },
    nameClasses() {
      return {
        "text-right": this.isRight,
      };
    },
    iconClasses() {
      return {
        "order-first": this.isLeft,
        "mr-3": this.isLeft,
        "ml-1": this.isLeft,
        "ml-3": this.isRight,
        "mr-1": this.isRight,
      };
    },
    slotBarClasses() {
      return {
        "order-first": this.isLeft,
        "slot-unknown": !this.hero,
        "slot-blue": this.hero && this.slot.index === 0,
        "slot-teal": this.hero && this.slot.index === 1,
        "slot-purple": this.hero && this.slot.index === 2,
        "slot-yellow": this.hero && this.slot.index === 3,
        "slot-orange": this.hero && this.slot.index === 4,
        "slot-pink": this.hero && this.slot.index === 5,
        "slot-olive": this.hero && this.slot.index === 6,
        "slot-light-blue": this.hero && this.slot.index === 7,
        "slot-green": this.hero && this.slot.index === 8,
        "slot-brown": this.hero && this.slot.index === 9,
        "left": this.isLeft,
        "right": this.isRight,
      };
    },
  },
};
</script>

<style lang="scss" scoped>
.live-match-player {
  width: 100%;
}

.slot-bar {
  width: 10px;
  height: 36px;

  &.left {
    margin-right: 4px;
  }

  &.right {
    margin-left: 4px;
  }
}

.slot-unknown {
  background-color: #333;
}

.slot-blue {
  background-color: #3074f9;
}

.slot-teal {
  background-color: #66ffc0;
}

.slot-purple {
  background-color: #bd00b7;
}

.slot-yellow {
  background-color: #f8f50a;
}

.slot-orange {
  background-color: #ff6901;
}

.slot-pink {
  background-color: #ff88c5;
}

.slot-olive {
  background-color: #a2b349;
}

.slot-light-blue {
  background-color: #63dafa;
}

.slot-green {
  background-color: #01831f;
}

.slot-brown {
  background-color: #9f6b00;
}
</style>
