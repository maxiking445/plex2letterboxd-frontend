<template>
    <div class="additional-link">
        <button :class="{ active: active }" @click="$emit('buttonClick')">
            <VueSpinner v-if="showAnimation" class="spinner" size="20" color="white" />
            <font-awesome-icon v-else :icon="currentIcon" :style="{ color: iconColor }" />
            {{ title }}
        </button>
    </div>
</template>

<script setup lang="ts">

import {
    VueSpinner,
} from 'vue3-spinners';
import { watch } from 'vue';
import { ref } from 'vue';

const props = defineProps<{
    title: string;
    icon: string;
    active: Boolean;
    showAnimation: Boolean;
}>();

const iconColor = ref("black");
const showAnimationInternal = ref(false);
const currentIcon = ref("");

showAnimationInternal.value = false;
currentIcon.value = props.icon;

watch(
    () => props.showAnimation,
    (newVal, oldVal) => {
        if (newVal) {
            showAnimationInternal.value = true;
            iconColor.value = "black"
        } else {
            showAnimationInternal.value = false;
            currentIcon.value = "check";
            iconColor.value = "limegreen";

            setTimeout(() => {
                currentIcon.value = props.icon;
                iconColor.value = "black"
            }, 1200);
        }
    },
    { immediate: false }
);

</script>

<style>
.additional-link button {
    gap: 5px;
    display: flex;
    align-items: center;
    padding: 0.3rem 0.8rem;
    font-weight: 600;
    border-radius: 0.25rem;
    border: none;
    background-color: #2d3138;
    color: #fff;
    cursor: pointer;

    transition: background-color 0.2s;
}

.additional-link button.active {
    background-color: #E5A00D;
    color: #181b20;
}

.additional-link button:hover {
    background-color: #44484f;
}

</style>