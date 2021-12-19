<template>
  <v-container>
    <v-form>
      <v-container>
        <v-row>
          <v-col cols="12" md="4">
            <v-text-field v-model="prod_cd" label="商品コード"></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field v-model="customer_cd" label="顧客コード"></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-text-field v-model="customer_name" label="顧客名"></v-text-field>
          </v-col>
        </v-row>
        <v-btn color="success" @click="submit">
          send!
        </v-btn>
      </v-container>
    </v-form>
    <div>
      <v-list width="auto">
        <v-list-item v-for="item in items">
          <v-card>
            <template>
              <v-list-item-content>
                <img v-bind:src="getImage(item)" height="300px">
                </img>
              </v-list-item-content>
            </template>
            <v-list-item-content>商品コード　{{item.prod_cd}}</v-list-item-content>
            <v-list-item-content>顧客コード　{{item.customer_cd}}</v-list-item-content>
            <v-list-item-content>顧客名　　　{{item.customer_name}}</v-list-item-content>
            <v-list-item-content>コメント　　{{item.comment}}</v-list-item-content>
          </v-card>
        </v-list-item>
      </v-list>
    </div>
  </v-container>
</template>

<script>
  import axios from 'axios'
  import Vuex from 'vuex'

  export default {
    data() {
      return {
        prod_cd: "",
        customer_cd: "",
        customer_name: "",
        items: "",
        imgSrc: "",
        headers: {
          'Content-Type': 'application/json;charset=UTF-8',
        },
      }
    },
    methods: {
      submit() {
        var params = new URLSearchParams()
        params.append('prod_cd', this.prod_cd)
        params.append('customer_cd', this.customer_cd)
        params.append('customer_name', this.customer_name)
        axios.post('/api/gazo', params, ).then((response) => {
            console.log(response.data);
            this.items = response.data
          })
          .catch(err => {
            if (err.response) {
              console.log(err.response);
            }
          })
      },
      getImage(item) {
        try {
          this.imgSrc = require('@/assets/image/' + item.img_name)
          return require('@/assets/image/' + item.img_name)
        } catch (error) {
          return ""
        }
      },
    },
  }

</script>
