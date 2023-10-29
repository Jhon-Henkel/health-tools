<template>
  <ion-page>
    <ion-header :translucent="true">
      <ion-toolbar>
        <ion-title>
          Pressão
        </ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <ion-content color="light">
        <ion-row class="ion-padding-top">
          <ion-col>
            <ion-datetime-button datetime="datetime" />
            <ion-modal :keep-contents-mounted="true">
              <ion-datetime presentation="month-year" id="datetime" />
            </ion-modal>
          </ion-col>
        </ion-row>
        <ion-row>
          <ion-col>
            <ion-list :inset="true">
              <ion-item>
                <ion-label>Sis/ Dia</ion-label>
                <ion-label>Bpm</ion-label>
                <ion-label>Data</ion-label>
                <ion-label>Hora</ion-label>
              </ion-item>
              <ion-item>
                <ion-label>110/ 11</ion-label>
                <ion-label>115</ion-label>
                <ion-label>25/10</ion-label>
                <ion-label>15:30</ion-label>
              </ion-item>
            </ion-list>
          </ion-col>
        </ion-row>
        <ion-row>
          <ion-col>
            <ion-button class="ion-margin-horizontal ion-margin-top" expand="block" href="/pressao/registrar">
              Inserir Novo
            </ion-button>
          </ion-col>
        </ion-row>
      </ion-content>
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
  import {
    IonContent,
    IonHeader,
    IonItem,
    IonToolbar,
    IonTitle,
    IonPage,
    IonList,
    IonButton,
    IonRow,
    IonCol,
    IonLabel,
    IonDatetimeButton,
    IonModal,
    IonDatetime
  } from '@ionic/vue';
  import { defineComponent } from 'vue';

  export default defineComponent({
    components: {
      IonContent,
      IonHeader,
      IonItem,
      IonToolbar,
      IonTitle,
      IonPage,
      IonList,
      IonButton,
      IonRow,
      IonCol,
      IonLabel,
      IonDatetimeButton,
      IonModal,
      IonDatetime
    }
  });
</script>

<style scoped>
  ion-datetime {
    border-radius: 16px;
  }
</style>