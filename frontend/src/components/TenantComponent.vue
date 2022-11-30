<template>
  <q-card>
    <div class="row">
      <div class="col-2 q-pa-md">
        <q-item>
          <q-item-section avatar>
            <q-icon color="primary" name="people"></q-icon>
          </q-item-section>

          <q-item-section class="text-h6 text-primary">
            Tenants
          </q-item-section>
        </q-item>
        <q-list>
          <q-item
            clickable
            v-for="tenant in selectedProperty.tenants"
            :key="tenant"
            :value="tenant"
            @click="tenantClick(tenant)"
            :style="
              tenant.subid === selectedTenant.subid
                ? { 'box-shadow': '0px 0px 3px 4px #6175cc' }
                : { background: 'none' }
            "
          >
            <q-item-section avatar>
              <q-icon color="primary" name="person"> </q-icon>
            </q-item-section>

            <q-item-section>
              <q-item-label> {{ tenant.name }} {{ tenant.unit }} </q-item-label>
            </q-item-section>
          </q-item>
        </q-list>
      </div>

      <!-- Tenant Dashboard Component -->
      <div class="col-9">
        <q-card>
          <q-item>
            <q-toolbar>
              <q-toolbar-title style="text-align: left"
                >Tenant Information</q-toolbar-title
              >

              <q-btn flat round dense icon="whatshot" />
            </q-toolbar>
          </q-item>
          <q-separator horizontal></q-separator>
          <q-item v-if="selectedTenant.subid != -1">
            <q-item>
              <q-item-section>
                <q-item-label>{{ selectedTenant.leasestart }}</q-item-label>
                <q-item-label caption>Lease Start Date</q-item-label>
              </q-item-section>
              <q-item-section>
                <q-item-label>{{ selectedTenant.currenttenant }}</q-item-label>
                <q-item-label caption>Currently Lives Here?</q-item-label>
              </q-item-section>
            </q-item>
            <q-item>
              <q-item-section>
                <q-item-label>{{ selectedTenant.leaseend }}</q-item-label>
                <q-item-label caption>Lease End Date</q-item-label>
              </q-item-section>
            </q-item>
          </q-item>
        </q-card>
      </div>
    </div>
  </q-card>
</template>

<script lang="ts">
import { usePortfolioStore } from 'src/stores/portfolio-store';
import { defineComponent, ref, computed } from 'vue';
import { Tenant, Expense, Revenue } from 'src/stores/portfolio-store';

export default defineComponent({
  name: 'TenantComponent',
  setup() {
    let portfolioStore = usePortfolioStore();

    const persistentTenants = ref(false);
    const selectedProperty = computed(
      () => portfolioStore.form.selectedProperty
    );

    function tenantClick(tenant: Tenant) {
      portfolioStore.form.selectedTenant = tenant;
    }

    const selectedTenant = computed(() => portfolioStore.form.selectedTenant);

    return {
      persistentTenants,
      selectedProperty,
      tenantClick,
      selectedTenant,
    };
  },
});
</script>
