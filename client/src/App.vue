<script setup lang="ts">
// imports
import axios from 'axios'
import { reactive, ref } from 'vue'
let appData = reactive([])
let databool = ref<boolean>(false)

async function getMeal() {
    appData.length = 0
    databool.value = !true
    try {
        let { data } = await axios.get('http://localhost:8080')
        console.log(data)
        // @ts-ignore
        appData.push(data)
        databool.value = true
    } catch (error) {
        console.log(error)
    }
}
</script>

<template>
    <div class="container">
        <section class="header">
            <h1>Feeling Hungry ?</h1>
            <div>Get A Random Meal</div>
            <button @click="getMeal">click for foods</button>
        </section>

        <section class="body" v-if="databool">
            <section class="pic-res">
                <div class="img">
                    <img
                        :src="appData[0].meals[0].strMealThumb"
                        alt="your food ig"
                    />
                </div>
                <span class="cat"
                    >Category:{{ appData[0].meals[0].strCategory }}</span
                >
                <br />
                <span class="cat">Area: {{ appData[0].meals[0].strArea }}</span>
                <div class="ings">
                    <h3 class="ingredients">Ingredients:</h3>
                    <ul>
                    </ul>
                </div>
            </section>
            <section class="tit-method"></section>
            {{ appData[0] }}
        </section>
    </div>
</template>

<style scoped lang="scss">
/*since it's a single fine All styles are in main.scss*/
</style>
