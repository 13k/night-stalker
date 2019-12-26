<template>
  <div class="live-match-player" :class="sideClass">
    <img
      class="player-hero-image"
      :src="player.hero | heroImageURL({ version: 'small' })"
      :title="player.hero | heroName('Picking hero')"
    />

    <router-link
      class="player-name"
      :class="slotColorClass"
      :to="{ name: 'players.show', params: { accountId: player.account_id } }"
    >
      {{ player.name }}
    </router-link>
  </div>
</template>

<script>
import filters from "@/components/filters";

export default {
  name: "live-match-player",
  filters,
  props: {
    player: Object,
    side: String
  },
  computed: {
    sideClass() {
      return {
        radiant: this.side !== "right",
        dire: this.side === "right"
      };
    },
    slotColorClass() {
      return {
        blue: this.player.player_slot === 0,
        teal: this.player.player_slot === 1,
        purple: this.player.player_slot === 2,
        yellow: this.player.player_slot === 3,
        orange: this.player.player_slot === 4,
        pink: this.player.player_slot === 5,
        olive: this.player.player_slot === 6,
        "light-blue": this.player.player_slot === 7,
        green: this.player.player_slot === 8,
        brown: this.player.player_slot === 9
      };
    }
  }
};
</script>

<style lang="scss" scoped>
.live-match-player {
  display: flex;
  align-items: center;
  margin: 3px 0 3px 0;
}

.live-match-player.radiant {
  flex-direction: row;
}

.live-match-player.dire {
  flex-direction: row-reverse;
}

.live-match-player.dire .player-name {
  text-align: right;
}

.player-name {
  flex-grow: 2;
  text-align: left;
  text-shadow: 1px 1px 1px #000;
  text-decoration: none;
}

.player-name.blue {
  color: #3074f9;
}

.player-name.teal {
  color: #66ffc0;
}

.player-name.purple {
  color: #bd00b7;
}

.player-name.yellow {
  color: #f8f50a;
}

.player-name.orange {
  color: #ff6901;
}

.player-name.pink {
  color: #ff88c5;
}

.player-name.olive {
  color: #a2b349;
}

.player-name.light-blue {
  color: #63dafa;
}

.player-name.green {
  color: #01831f;
}

.player-name.brown {
  color: #9f6b00;
}

.player-hero-image {
  width: 59px;
  height: 33px;
  margin-right: 8px;
  margin-left: 8px;
  box-shadow: 2px 2px 4px #000, -1px -1px 2px #333;
}
</style>
