<template>
  <ion-page>
    <ion-tabs>
      <ion-router-outlet />
      <ion-tab-bar slot="bottom">
        <ion-tab-button tab="blood-pressure" href="/pressao">
          <ion-icon :icon="heartCircleOutline" />
          <ion-label>Pressão</ion-label>
        </ion-tab-button>
        <ion-tab-button tab="blood-glucose" href="/glicemia">
          <ion-icon :icon="waterOutline" />
          <ion-label>Glicemia</ion-label>
        </ion-tab-button>
      </ion-tab-bar>
    </ion-tabs>
  </ion-page>
</template>

<script lang="ts">
  import { IonPage, IonTabs, IonRouterOutlet, IonTabBar, IonTabButton, IonLabel, IonIcon } from '@ionic/vue';
  import { heartCircleOutline, waterOutline } from 'ionicons/icons';

  export default {
    components: { IonPage, IonTabs, IonRouterOutlet, IonTabBar, IonTabButton, IonLabel, IonIcon },
    data() {
      return {
        heartCircleOutline,
        waterOutline,
      };
    },
  };
</script>