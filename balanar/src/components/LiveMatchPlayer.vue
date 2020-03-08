<template>
  <div
    :class="containerClasses"
    class="d-flex align-center flex-grow-1"
  >
    <div
      :class="nameClasses"
      class="d-flex flex-column"
    >
      <span>{{ name }}</span>

      <span
        v-if="kda"
        class="overline grey--text"
      >
        {{ kda }}
      </span>
    </div>

    <div
      :class="iconClasses"
      class="d-flex"
    >
      <HeroImage
        :hero="hero"
        orientation="landscape"
        width="64"
        height="36"
        alt-placeholder="Picking hero"
        alt
        placeholder
      />

      <div
        :class="slotBarClasses"
        class="slot-bar"
      />
    </div>
  </div>
</template>

<script>
import * as $t from "@/protocol/transform";
import pb from "@/protocol/proto";
import { normalizePlayerName } from "@/util";
import HeroImage from "@/components/HeroImage.vue";

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
      type: pb.protocol.LiveMatch.Player,
      required: true,
    },
  },

  computed: {
    hero() {
      return $t.get(this.player, "hero");
    },
    slot() {
      return $t.get(this.player, "slot");
    },
    label() {
      const label = normalizePlayerName(this.player.label);
      const name = normalizePlayerName(this.player.name);

      if (name !== label) {
        return this.player.label;
      }

      return null;
    },
    name() {
      let name = this.player.name;

      if (this.label) {
        name = `${name} (${this.label})`;
      }

      return name;
    },
    kda() {
      const { kills = 0, deaths = 0, assists = 0 } = this.player || {};

      if (kills > 0 || deaths > 0 || assists > 0) {
        return `${kills}/${deaths}/${assists}`;
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
        "align-end": this.isRight,
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
