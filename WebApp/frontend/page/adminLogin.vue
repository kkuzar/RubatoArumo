<template>
    <v-app id="inspire">
        <v-content>
            <v-container fluid fill-height>
                <v-layout align-center justify-center>
                    <v-flex xs12 sm8 md4>
                        <v-card class="elevation-12">
                            <v-img
                                    src="/public/img/sf-logo_s.png"
                                    aspect-ratio="4.1"
                            ></v-img>
                            <v-card-text>
                                <v-form>
                                    <v-text-field prepend-icon="person" name="login" label="Login" type="text"></v-text-field>
                                    <v-text-field id="password" prepend-icon="lock" name="password" label="Password" type="password"></v-text-field>
                                </v-form>
                            </v-card-text>
                            <v-card-actions>
                                <v-spacer></v-spacer>
                                <v-btn color="primary">Login</v-btn>
                            </v-card-actions>
                        </v-card>
                    </v-flex>
                </v-layout>
            </v-container>
        </v-content>
    </v-app>
</template>

<script>
export default {
    name: "adminLogin"
}
</script>

<style scoped>
#inspire {
    background: url("/public/img/adminLogin_bk.jpg") no-repeat center center fixed;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
}
</style>