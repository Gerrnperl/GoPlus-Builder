<template>
  <div class="ui-empty" :class="[`size-${size}`]">
    <template v-if="size === 'large'">
      <img :src="emptyImg" />
      <slot></slot>
      <template v-if="!slots.default">
        {{ defaultText }}
      </template>
    </template>
    <template v-else-if="size === 'extra-large'">
      <img :src="emptyXlImg" />
      <p class="text">
        <slot></slot>
        <template v-if="!slots.default">
          {{ defaultText }}
        </template>
      </p>
      <div class="op"><slot name="op"></slot></div>
    </template>
    <template v-else>
      <svg
        v-if="size === 'medium'"
        class="line"
        width="36"
        height="2"
        viewBox="0 0 36 2"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M1 1H9"
          stroke="#EAEFF3"
          stroke-width="1.56"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
        <path
          d="M15 1H35"
          stroke="#EAEFF3"
          stroke-width="1.56"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
      <slot></slot>
      <template v-if="!slots.default">
        {{ defaultText }}
      </template>
      <svg
        v-if="size === 'medium'"
        class="line"
        width="36"
        height="2"
        viewBox="0 0 36 2"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M1 1H21"
          stroke="#EAEFF3"
          stroke-width="1.56"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
        <path
          d="M27 1H35"
          stroke="#EAEFF3"
          stroke-width="1.56"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
    </template>
  </div>
</template>

<script setup lang="ts">
import { computed, useSlots } from 'vue'
import { useConfig } from '../UIConfigProvider.vue'
import emptyImg from './empty.svg'
import emptyXlImg from './empty-xl.svg'

defineProps<{
  size: 'small' | 'medium' | 'large' | 'extra-large'
}>()

const config = useConfig()
const defaultText = computed(() => config.empty?.text ?? 'No data')
const slots = useSlots()

// TODO: support button for size:large ?
</script>

<style scoped lang="scss">
.ui-empty {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.size-large {
  flex-direction: column;
  gap: 12px;

  font-size: 16px;
  line-height: 26px;
  color: var(--ui-color-grey-1000);
}

.size-extra-large {
  flex-direction: column;

  .text {
    margin-top: 4px;
    color: var(--ui-color-grey-700);
    font-size: 16px;
    line-height: 26px;
  }

  .op {
    margin-top: 24px;
    display: flex;
    gap: var(--ui-gap-large);
  }
}

.size-small,
.size-medium {
  color: var(--ui-color-grey-600);
  gap: 8px;
}
</style>
