<template>
  <div
    class="live-match-player d-flex align-center"
    :class="containerClasses"
  >
    <span :class="nameClasses">
      {{ player.name }}
    </span>

    <div
      class="icon d-flex"
      :class="iconClasses"
    >
      <HeroImage
        :hero="player.hero"
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
      type: Object,
      required: true,
    },
  },

  computed: {
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
        "mr-2": this.isLeft,
        "ml-1": this.isLeft,
        "ml-2": this.isRight,
        "mr-1": this.isRight,
      };
    },
    slotBarClasses() {
      return {
        "order-first": this.isLeft,
        "slot-blue": this.player.player_slot === 0,
        "slot-teal": this.player.player_slot === 1,
        "slot-purple": this.player.player_slot === 2,
        "slot-yellow": this.player.player_slot === 3,
        "slot-orange": this.player.player_slot === 4,
        "slot-pink": this.player.player_slot === 5,
        "slot-olive": this.player.player_slot === 6,
        "slot-light-blue": this.player.player_slot === 7,
        "slot-green": this.player.player_slot === 8,
        "slot-brown": this.player.player_slot === 9,
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

.icon {
  box-shadow: 2px 2px 4px #000, -1px -1px 2px #333;
}

.slot-bar {
  width: 6px;
  height: 36px;

  &.left {
    border-right: 1px solid #444;
  }

  &.right {
    border-left: 1px solid #444;
  }
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
