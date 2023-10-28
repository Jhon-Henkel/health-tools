<template>
  <ion-page>
    <ion-header :translucent="true">
      <ion-toolbar>
        <ion-title>
          Registro de pressão
        </ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <ion-content color="light">
        <ion-row>
          <ion-col>
              <ion-list :inset="true">
                <ion-item>
                  <ion-input label="Sistólica" type="number"/>
                </ion-item>
                <ion-item>
                  <ion-input label="Diastólica" type="number"/>
                </ion-item>
                <ion-item>
                  <ion-input label="Batimentos" type="number"/>
                </ion-item>
              </ion-list>
              <ion-note color="medium" class="ion-margin-horizontal">
                Preencha com os dados da sua última aferição.
              </ion-note>
              <ion-button class="ion-margin-horizontal ion-margin-top" expand="block">
                Salvar
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
    IonInput,
    IonItem,
    IonNote,
    IonToolbar,
    IonTitle,
    IonPage,
    IonMenuButton,
    IonButtons,
    IonList,
    IonButton,
    IonRow,
    IonCol
  } from '@ionic/vue';
  import { defineComponent } from 'vue';

  export default defineComponent({
    components: {
      IonContent,
      IonHeader,
      IonInput,
      IonItem,
      IonNote,
      IonToolbar,
      IonTitle,
      IonPage,
      IonMenuButton,
      IonButtons,
      IonList,
      IonButton,
      IonRow,
      IonCol
    }
  });
</script>