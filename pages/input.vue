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
          <v-col cols="12" md="4">
            <v-text-field v-model="comment" label="コメント"></v-text-field>
          </v-col>
          <v-col cols="12" md="4">
            <v-file-input v-model="img_file" accept="image/*" label="画像を選択して下さい" show-size @change="onChangeFileInput">
            </v-file-input>
          </v-col>
        </v-row>
        <v-btn color="success" @click="submit">
          send!
        </v-btn>
      </v-container>
    </v-form>
    <div>
      <v-list>
        <v-list-item v-for="item in info">
          <v-card>
            <v-list-item-content>
              <v-img :src="require('@/assets/image/' + item.img_name)" height="300px">
              </v-img>
            </v-list-item-content>
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
  import axios from 'axios';
  import Vuex from 'vuex';
  //import FormData from "form-data";

  export default {
    data() {
      return {
        prod_cd: "",
        customer_cd: "",
        customer_name: "",
        comment: "",
        img_file: "",
        formData: null,
      }
    },
    methods: {


      async uploadImageFile() {
        let formData = new FormData()
        formData.append("file", this.img_file);
        const config = {
          headers: {
            "Content-type": "multipart/form-data",
          },
        };

        const res = await axios.post('/api/upload', formData, config)
          .then((response) => {
            console.log("upload successed");
            console.log(this.img_file.name)
          })
          .catch((e) => {
            console.log("upload failed")
            alert("ファイルのアップロードが出来ませんでした。")
          })
      },

      submit() {
        var params = new URLSearchParams()
        params.append('prod_cd', this.prod_cd)
        params.append('customer_cd', this.customer_cd)
        params.append('customer_name', this.customer_name)
        params.append('comment', this.comment)
        params.append('img_name', this.img_file.name)
        if (this.img_file !== "") {
          this.uploadImageFile();
        }
        axios.post('/api/create_gazo', params, ).then((response) => {
            console.log(response.data);
            this.prod_cd = ""
            this.customer_cd = ""
            this.customer_name = ""
            this.comment = ""
            this.img_file = ""
          })
          .catch(err => {
            if (err.response) {
              console.log(err.response);
            }
          })
      },
    },
  }

</script>
