<template>
  <q-card v-ripple class="my-card">
    <q-card-section class="bg-primary text-white">
      <div class="text-h6">property_summary.description</div>
      <div class="text-subtitle2">property_summary.type</div>
    </q-card-section>
  </q-card>
</template>

<script lang="ts">
import { PropertySummary } from './models';
import { defineComponent } from 'vue';

type propertyprops = {
  propertySummary: PropertySummary;
};

export default defineComponent({
  name: 'PropertyCard',
  props: { propertySummary: { type: PropertySummary, required: true } },
  setup(props: propertyprops) {
    const property_summary = props.propertySummary;

    return {
      property_summary,
    };
  },
});
</script>

<style lang="sass" scoped>
.my-card
  width: 100%
  height: 100%
  max-width: 18vw
  margin-left: 3vw
</style>
