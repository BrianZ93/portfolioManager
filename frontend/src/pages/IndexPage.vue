<template>
  <q-page>
    <IndexComponent title="Daily Pulse"></IndexComponent>
  </q-page>
</template>

<script>
import { defineComponent } from 'vue';
import IndexComponent from 'src/components/IndexComponent.vue';

export default defineComponent({
  name: 'IndexPage',
  components: { IndexComponent },
  setup() {
    const indexes = [
      {
        value: 1,
      },
      {
        value: 2,
      },
      {
        value: 3,
      },
      {
        value: 4,
      },
    ];
    return { indexes };
  },
});
</script>
