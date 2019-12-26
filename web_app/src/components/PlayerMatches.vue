<template>
  <div class="player-matches">
    <ul class="matches">
      <li
        class="match"
        v-for="match in matches"
        :key="match.match_id"
        :match="match"
      >
        <img
          class="hero-image"
          :src="
            match.hero | heroImageURL({ version: 'small', placeholder: true })
          "
          :title="match.hero | heroName"
        />

        <a
          class="match-id"
          :href="`https://www.opendota.com/matches/${match.match_id}`"
          :title="`View match ${match.match_id} on OpenDota`"
          target="_blank"
        >
          {{ match.match_id }}
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
import filters from "@/components/filters";

export default {
  name: "player-matches",
  filters,
  props: {
    matches: Array
  }
};
</script>

<style lang="scss" scoped>
.matches {
  padding: 0;
  margin: 2em 0 0 2em;
}

.match {
  display: flex;
  align-items: center;
  padding: 0;
  margin: 10px 0;
  list-style: none;
}

.hero-image {
  width: 59px;
  height: 33px;
  margin-right: 8px;
  margin-left: 8px;
  box-shadow: 2px 2px 4px #000, -1px -1px 2px #333;
}

.match-id {
  margin-left: 16px;
  font-weight: bold;
}
</style>
