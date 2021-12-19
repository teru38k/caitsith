<template>
  <v-container>
    <v-layout>
      <v-dialog v-model="searchDialog" persistent width="600px">
        <v-btn class="mr-auto" @click="searchDialog = true" color="primary" dark slot="activator">登録検索</v-btn>
        <v-card>
          <v-card-title>
            <span class="headline">登録検索</span>
          </v-card-title>
          <v-card-text>
            <v-container grid-list-md>
              <v-layout wrap>
                <v-flex xs12 sm6 md4>
                  <v-text-field v-model="sProd_cd" label="商品コード"></v-text-field>
                </v-flex>
                <v-flex xs12 sm6 md4>
                  <v-text-field v-model="sCustomer_cd" label="顧客コード"></v-text-field>
                </v-flex>
                <v-flex xs12 sm6 md4>
                  <v-text-field v-model="sCustomer_name" label="顧客名"></v-text-field>
                </v-flex>
              </v-layout>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" flat @click.native="closeSerchDialog(); searchDialog = false">Close</v-btn>
            <v-btn color="primary" flat @click.native="getList(); searchDialog = false">Search</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-dialog v-model="inputDialog">
        <v-btn class="ml-auto" @click="inputDialog = true" color="primary" dark slot="activator">登録</v-btn>
        <v-card>
          <v-card-title>
            <span class="headline">登録</span>
          </v-card-title>
          <v-card-text>
            <v-container grid-list-md>
              <v-row>
                <v-col cols="12" md="4">
                  <v-text-field v-model="prod_cd" label="商品コード"></v-text-field>
                </v-col>
              </v-row>
              <v-row>
                <v-col cols="12" md="4">
                  <v-text-field v-model="customer_cd" label="顧客コード"></v-text-field>
                </v-col>
                <v-col cols="12" md="8">
                  <v-text-field v-model="customer_name" label="顧客名"></v-text-field>
                </v-col>
              </v-row>
              <v-col>
                <v-textarea v-model="comment" label="コメント"></v-textarea>
              </v-col>
              <v-row>
                <v-col cols="12" md="4">
                  <v-file-input v-model="img_file" accept="image/*" label="画像を選択して下さい" show-size
                    @change="onChangeFileInput">
                  </v-file-input>
                </v-col>
              </v-row>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn color="primary" flat @click.native="closeInputDialog(); inputDialog = false">Close</v-btn>
            <v-btn color="primary" flat @click.native="createItem(); inputDialog = false">Input!</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </v-layout>
    <br>

    <div>
      <v-data-table :headers="headers" :items="items" hide-actions class="elevation-1" select-all
        :pagination.sync="pagination" :item-key="id" :search="search">
        <template v-slot:item.actions="{ item }">
          <v-icon medium class="mr-2" @click="editDialog = true; getItem(item); getImage(item);">
            mdi-pencil
          </v-icon>
          <v-icon medium class="mr-2" @click="OpenDeleteDialog(item);">
            mdi-delete
          </v-icon>
        </template>
      </v-data-table>
      <v-layout row justify-center>
        <v-dialog v-model="editDialog" persistent>
          <v-card>
            <v-card-title>
              <span class="headline">登録内容編集</span>
            </v-card-title>
            <template>
              <div>
                <img :src="imgSrc" height="150px">
                </img>
              </div>
            </template>
            <v-card-text>
              <v-container>
                <v-row>
                  <v-text-field v-model="prod_cd" label="商品コード" required></v-text-field>
                </v-row>
                <v-row>
                  <v-col cols="12" md="4">
                    <v-text-field v-model="customer_cd" label="顧客コード" required></v-text-field>
                  </v-col>
                  <v-col cols="12" md="8">
                    <v-text-field v-model="customer_name" label="顧客名" required></v-text-field>
                  </v-col>
                </v-row>
                <v-row>
                  <v-col>
                    <v-textarea v-model="comment" label="コメント"></v-textarea>
                  </v-col>
                </v-row>
                <v-row>
                  <v-layout>
                    <v-dialog v-model="imageSelect" scrollable max-width="300px">
                      <v-btn class="ma-2" color="secondary" dark slot="activator" @click="getImages()">
                        画像選択
                        <v-icon>mdi-database-search-outline</v-icon>
                      </v-btn>
                      <v-card>
                        <v-card-subtitle>サーバー内の画像を選択してください。</v-card-subtitle>
                        <v-divider></v-divider>
                        <v-card-text>選択中の画像: {{img_name}}</v-card-text>
                        <v-card-text style="height: 350px;">
                          <div>
                        <vue-select-image :data-images="images" @onselectimage="onSelectImage" :h="70">
                        </vue-select-image>
                      </div>
                        </v-card-text>
                        <v-divider></v-divider>
                        <v-card-actions>
                          <v-btn color="primary" flat @click.native="closeImageSelect()">Close</v-btn>
                        </v-card-actions>
                      </v-card>
                    </v-dialog>
                    <p justify-center>選択中の画像: {{img_name}} </p>
                  </v-layout>
                </v-row>
                <br>
                <br>
                <v-row>
                  <v-file-input v-model="img_file" accept="image/*" label="アップロードする画像を選択して下さい" show-size
                    @change="onChangeFileInput">
                  </v-file-input>
                </v-row>
              </v-container>
              <small>*indicates required field</small>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="primary" flat @click.native="closeInputDialog(); editDialog = false">Close</v-btn>
              <v-btn color="primary" flat @click.native="updateItem(); editDialog = false">Save</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-layout>
      <v-dialog
        v-model="deleteDialog" 
        persistent :overlay="false"
        max-width="300px"
      >
        <v-card>
          <v-card-title>削除確認</v-card-title>
          <v-card-text>ID: {{id}} を削除しますか？</v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn class="ml-auto" color="info" @click="closeDeleteDialog()">キャンセル</v-btn>
            <v-btn class="mr-auto" color="warning" @click="deleteConfirm()">削除</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </v-container>
</template>

<script>
  import axios from 'axios'
  import Vuex from 'vuex'

  export default {
    data() {
      return {
        searchDialog: false,
        inputDialog: false,
        editDialog: false,
        deleteDialog: false,
        imageSelect: false,
        id: "",
        prod_cd: "",
        customer_cd: "",
        customer_name: "",
        comment: "",
        img_name: "",
        img_file: "",
        sProd_cd: "",
        sCustomer_cd: "",
        sCustomer_name: "",
        s2Prod_cd: "",
        s2Customer_cd: "",
        s2Customer_name: "",
        search: "",
        pagination: "30",
        headers: [{
            text: 'ID',
            value: 'id'
          },
          {
            text: '商品コード',
            value: 'prod_cd'
          },
          {
            text: '得意先コード',
            value: 'customer_cd'
          },
          {
            text: '得意先名',
            value: 'customer_name'
          },
          {
            text: 'コメント',
            value: 'comment'
          },
          {
            text: '編集/削除',
            value: 'actions',
            sortable: false
          }
        ],
        items: [],
        imgSrc: "",
        images: [],
      }
    },
    methods: {
      closeSerchDialog() {
        this.sProd_cd = ""
        this.sCustomer_cd = ""
        this.sCustomer_name = ""
      },
      closeInputDialog() {
        this.prod_cd = ""
        this.customer_cd = ""
        this.customer_name = ""
        this.comment = ""
        this.img_file = ""
        this.img_name = ""
      },
      closeImageSelect() {
        this.images = []
        this.imageSelect = false
      },
      closeDeleteDialog() {
        this.id = ""
        this.deleteDialog = false
      },
      OpenDeleteDialog(item) {
        this.deleteDialog = true
        this.id = item.id
      },

      getList() {
        var params = new URLSearchParams()
        params.append('prod_cd', this.sProd_cd)
        params.append('customer_cd', this.sCustomer_cd)
        params.append('customer_name', this.sCustomer_name)
        axios.post('/api/gazo', params, ).then((response) => {
            console.log(response.data);
            if (response.data === null) {
              alert("検索結果がありません。")
              this.items = []
              this.sProd_cd = ""
              this.sCustomer_cd = ""
              this.sCustomer_name = ""
            } else {
              this.items = response.data
              this.s2Prod_cd = this.sProd_cd
              this.s2Customer_cd = this.sCustomer_cd
              this.s2Customer_name = this.sCustomer_name
              this.sProd_cd = ""
              this.sCustomer_cd = ""
              this.sCustomer_name = ""
            }
          })
          .catch(err => {
            if (err.response) {
              console.log(err.response);
              alert("検索エラー")
              this.items = []
              this.sProd_cd = ""
              this.sCustomer_cd = ""
              this.sCustomer_name = ""
            }
          })
      },

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
            this.img_name = this.img_file.name
            console.log(this.img_name)
          })
          .catch((e) => {
            console.log("upload failed")
            alert("ファイルのアップロードが出来ませんでした。")
          })
      },

      createItem() {
        var params = new URLSearchParams()
        params.append('prod_cd', this.prod_cd)
        params.append('customer_cd', this.customer_cd)
        params.append('customer_name', this.customer_name)
        params.append('comment', this.comment)
        params.append('img_name', this.img_name)
        if (this.img_file !== "") {
          this.uploadImageFile();
        }
        axios.post('/api/create_gazo', params, ).then((response) => {
            console.log(response.data);
            alert("正常に追加されました。")
            this.prod_cd = ""
            this.customer_cd = ""
            this.customer_name = ""
            this.comment = ""
            this.img_file = ""
          })
          .catch(err => {
            if (err.response) {
              console.log(err.response);
              alert("正常に追加できませんでした。")
              this.prod_cd = ""
              this.customer_cd = ""
              this.customer_name = ""
              this.comment = ""
              this.img_file = ""
            }
          })
      },

      updateItem() {
        var params = new URLSearchParams()
        params.append('id', this.id)
        params.append('prod_cd', this.prod_cd)
        params.append('customer_cd', this.customer_cd)
        params.append('customer_name', this.customer_name)
        params.append('comment', this.comment)
        params.append('img_name', this.img_name)
        if (this.img_file !== "") {
          this.uploadImageFile();
        }
        axios.post('/api/update_gazo', params, ).then((response) => {
            console.log(response.data);
            alert("正常に更新されました。")
            this.id = ""
            this.prod_cd = ""
            this.customer_cd = ""
            this.customer_name = ""
            this.comment = ""
            this.img_file = ""
          })
          .catch(err => {
            if (err.response) {
              console.log(err.response);
              alert("正常に更新できませんでした。")
              this.id = ""
              this.prod_cd = ""
              this.customer_cd = ""
              this.customer_name = ""
              this.comment = ""
              this.img_file = ""
            }
          })
      },

      getImage(item) {
        try {
          this.imgSrc = require('@/assets/image/' + item.img_name)
          //return require('@/assets/image/' + item.img_name)
        } catch (error) {
          //return ""
          this.imgSrc = ""
        }
      },

      getImages() {
        this.imageSelect = true;
        var params = new URLSearchParams()
        axios.post('/api/getimages', params, ).then((response) => {
            console.log(response.data);
            if (response.data === null) {
              alert("検索結果がありません。")
              this.images = []
            } else {
              for (const elem of response.data) {
                console.log(elem);
                const elemImgName = elem
                const src = require('@/assets/image/' + elem)
                const alt = "Alt Image" + elem
                const dat = {
                  "img_name": elemImgName,
                  "src": src,
                  "alt": alt,
                }
                this.images.push(dat);
              };
            }
          })
          .catch(err => {
            if (err.response) {
              console.log(err.response);
              alert("検索エラー")
              this.images = []
            }
          })
      },

      onSelectImage: function (selected) {
        this.img_name = selected.img_name
      },

      getItem(item) {
        var params = new URLSearchParams()
        params.append('id', item.id)
        axios.post('/api/gazo', params, ).then((response) => {
            console.log(response.data);
            this.id = response.data[0]["id"]
            this.prod_cd = response.data[0]["prod_cd"]
            this.customer_cd = response.data[0]["customer_cd"]
            this.customer_name = response.data[0]["customer_name"]
            this.comment = response.data[0]["comment"]
            this.editDialog = true
          })
          .catch(err => {
            if (err.response) {
              console.log(err.response);
            }
          })
      },

      deleteConfirm() {
        var params = new URLSearchParams()
        params.append('id', this.id)
        axios.post('/api/delete_gazo', params, ).then((response) => {
          console.log(response.data);
          alert("ID: " + this.id + " を削除しました。")
          this.sProd_cd = this.s2Prod_cd
          this.sCustomer_cd = this.s2Customer_cd
          this.sCustomer_name = this.s2Customer_name
          this.getList()
          this.closeDeleteDialog()
        }).catch(err => {
          if (err.response) {
            console.log(err.response);
            alert("削除できませんでした。")
            this.closeDeleteDialog()
          }
        })
      },
    },
  }

</script>
