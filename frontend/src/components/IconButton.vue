<template>
    <div class="icon-btn-wrapper">
        <button :class="buttonClasses" @click="$emit('buttonClick')">
            <VueSpinner v-if="showAnimation" class="spinner" size="16" color="white" />
            <font-awesome-icon v-else :icon="currentIcon" :style="{ color: iconColor }" />
            {{ title }}
        </button>
    </div>
</template>

<script setup lang="ts">

import {
    VueSpinner,
} from 'vue3-spinners';
import { computed, watch } from 'vue';
import { ref } from 'vue';

const props = defineProps<{
    title: string;
    icon?: string;
    active: Boolean;
    showAnimation?: Boolean;
    status?: 'ready' | 'required' | 'checking';
}>();

const iconColor = ref("currentColor");
const showAnimationInternal = ref(false);
const currentIcon = ref("");

showAnimationInternal.value = false;
currentIcon.value = props.icon;

const buttonClasses = computed(() => ({
    active: props.active,
    ready: props.status === 'ready',
    required: props.status === 'required',
    checking: props.status === 'checking',
}));

watch(
    () => props.icon,
    (newVal) => {
        currentIcon.value = newVal || "";
        iconColor.value = "currentColor";
    },
);

watch(
    () => props.showAnimation,
    (newVal, oldVal) => {
        if (newVal) {
            showAnimationInternal.value = true;
            iconColor.value = "currentColor"
        } else {
            showAnimationInternal.value = false;
            currentIcon.value = "check";
            iconColor.value = "#22c55e";

            setTimeout(() => {
                currentIcon.value = props.icon;
                iconColor.value = "currentColor"
            }, 1200);
        }
    },
    { immediate: false }
);

</script>

<style scoped>
.icon-btn-wrapper button {
    gap: 6px;
    display: flex;
    align-items: center;
    padding: 0.4rem 0.85rem;
    font-size: 0.8rem;
    font-weight: 600;
    font-family: inherit;
    border-radius: var(--radius-sm);
    border: 1px solid transparent;
    background-color: var(--bg-surface);
    color: var(--text-secondary);
    cursor: pointer;
    transition: all var(--transition-fast);
    white-space: nowrap;
}

.icon-btn-wrapper button.active {
    background-color: var(--plex-accent);
    color: var(--bg-primary);
    border-color: transparent;
}

.icon-btn-wrapper button:hover {
    background-color: var(--bg-surface-hover);
    color: var(--text-primary);
    border-color: var(--border-medium);
}

.icon-btn-wrapper button.active:hover {
    background-color: var(--plex-accent-hover);
    color: var(--bg-primary);
    border-color: transparent;
}

.icon-btn-wrapper button.ready {
    border-color: rgba(34, 197, 94, 0.55);
    color: #22c55e;
}

.icon-btn-wrapper button.ready:hover {
    border-color: rgba(34, 197, 94, 0.8);
    color: #22c55e;
}

.icon-btn-wrapper button.required {
    border-color: rgba(239, 68, 68, 0.55);
    color: #ef4444;
}

.icon-btn-wrapper button.required:hover {
    border-color: rgba(239, 68, 68, 0.8);
    color: #ef4444;
}

.icon-btn-wrapper button.checking {
    border-color: rgba(229, 160, 13, 0.55);
    color: var(--plex-accent);
}

.icon-btn-wrapper button.checking svg {
    animation: icon-btn-spin 0.9s linear infinite;
}

@keyframes icon-btn-spin {
    to {
        transform: rotate(360deg);
    }
}
</style>
