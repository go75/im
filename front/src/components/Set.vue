<script lang="ts" setup>

import {ref} from 'vue'

const emit = defineEmits(['page', "msg"])
function emitPage(s:number) {
  emit('page', s)
}
</script>

<template>
    <div>

    </div>
</template>

<style scoped>

</style>
