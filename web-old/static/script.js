// Sandwich-Daemon script.js by TehRockettek
// https://github.com/TheRockettek/Sandwich-Daemon

// Install any plugins
Vue.use(VueChartJs);

Vue.component("line-chart", {
    extends: VueChartJs.Line,
    mixins: [VueChartJs.mixins.reactiveProp],
    props: ['chartData', 'options'],
    mounted() {
        this.renderChart(this.chartData, this.options)
    },
})

Vue.component("status-graph", {
    props: ['value', 'colours'],
    template: `
    <div class="progress">
        <div v-for="(value, index) in this.keys" :class="'progress-bar bg-' + colours[index]" role="progressbar" :style="'width: ' + (value/total)*100 + '%'" :aria-valuenow="(value/total)*100" aria-valuemin="0" aria-valuemax="100"></div>
    </div>
    `,
    data() {
        return {
            keys: {},
            total: 0,
        }
    },
    watch: {
        value: function () {
            this.loadValues()
        }
    },
    mounted() {
        this.loadValues()
    },
    methods: {
        loadValues() {
            this.keys = {}
            this.total = 0
            shards = Object.values(this.value.shards)
            for (shindex in shards) {
                shard = shards[shindex]
                if (shard.status in this.keys) {
                    this.keys[shard.status]++
                } else {
                    this.keys[shard.status] = 1
                }
                this.total++
            }
            this.$forceUpdate()
        }
    }
})

Vue.component("card-display", {
    props: ['title', 'value', 'bg'],
    template: `
    <div class="col justify-content-center d-flex">
        <div :class="bg+' card text-white m-1'" style="width: 18rem;">
            <div class="card-header">{{ title }} </div>
            <div class="card-body">
                <h5 class="card-title">{{ value }}</h5>
            </div>
        </div>
    </div>
    `,
})

Vue.component("form-submit", {
    props: {
        label: {
            default: "Save Changes",
        }
    },
    template: `<button type="submit" class="btn btn-dark" @click="onClick">{{ label }}</button>`,
    methods: {
        onClick() {
            this.$emit('click');
        }
    }
})

Vue.component("form-input", {
    props: ['type', 'id', 'label', 'values', 'value', 'disabled', 'placeholder'],
    template: `
    <div class="form-check" v-if="type == 'checkbox'">
        <input class="form-check-input" type="checkbox" :id="id" :checked="value" v-on:change="updateValue($event.target.checked)" :disabled="disabled">
        <label class="form-check-label" :for="id">{{ label }}</label>
    </div>
    <div class="mb-3" v-else-if="type == 'text'">
        <label :for="id" class="col-sm-12 form-label">{{ label }}</label>
        <input type="text" class="form-control" :id="id" :value="value" v-on:change="updateValue($event.target.value)" :disabled="disabled" :placeholder="placeholder">
    </div>
    <div class="mb-3" v-else-if="type == 'list'">
        <label :for="id" class="col-sm-12 form-label">{{ label }}</label>
        <input type="text" class="form-control" :id="id" :value="value" placeholder="A,B,C" v-on:change="updateValue(Array.from(new Set($event.target.value.split(','))))    " :disabled="disabled">
    </div>
    <div class="mb-3" v-else-if="type == 'number'">
        <label :for="id" class="col-sm-12 form-label">{{ label }}</label>
        <input type="number" class="form-control" :id="id" :value="value" v-on:change="updateValue(Number($event.target.value))" :disabled="disabled" :placeholder="placeholder">
    </div>
    <div class="mb-3" v-else-if="type == 'password'">
        <label :for="id" class="col-sm-12 form-label">{{ label }}</label>
        <div class="input-group">
            <input type="password" class="form-control" :id="id" autocomplete :value="value" v-on:change="updateValue($event.target.value)" :disabled="disabled" :placeholder="placeholder">
            <button class="btn btn-outline-dark" type="button" v-on:click="copyFormInputPassword()">Copy</button>
        </div>
    </div>
    <div class="mb-3" v-else-if="type == 'select'">
        <label :for="id" class="col-sm-12 form-label">{{ label }}</label>
        <select class="form-select" :id="id" v-on:change="updateValue($event.target.value)" :disabled="disabled">
            <option v-for="item in values" selected="item == value">{{ item }}</option>
        </select>
    </div>
    <div class="mb-3 row pb-4" v-else-if="type == 'intent'">
        <label for="managerBotIntents" class="col-sm-3 form-label">{{ label }}</label>
        <div class="col-sm-9">
            <input type="number" class="form-control" min=0 max=32767 :value="value" @change="(v) => {v.target.value = v.target.value & 32767; updateValue(Number(v.target.value)); fromIntents(v.target.value)}" @input="(v) => {updateValue(Number(v.target.value)); fromIntents(v.target.value)}" :disabled="disabled">
            <div class="form-row py-2">
                <div class="form-check form-check-inline col-sm-8 col-md-5" v-for="(intent, index) in this.intents">
                    <input class="form-check-input" type="checkbox" v-bind:value="index" v-bind:id="id+index" v-model="selectedIntent" @change="calculateIntent()">
                    <label class="form-check-label" v-bind:for="'managerBotIntentBox'+index">{{intent}}</label>
                </div>
            </div>
        </div>
    </div>
    <div class="mb-3 row pb-4" v-else-if="type == 'presence'">
        <label class="col-sm-3 col-form-label">{{ label }}</label>
        <div class="col-sm-9">
            <div class="mb-3">
                <label :for="id + 'status'" class="col-sm-12 form-label">Status</label>
                <select class="form-select" :id="id + 'status'" :value="value.status" @input="(v) => {value.status = v.target.value}">
                    <option v-for="item in ['','online','dnd','idle','invisible','offline']" :key="item" :disabled="!item" :selected="item == value">{{ item }}</option>
                </select>
            </div>
            <div class="mb-3">
                <label :for="id + 'name'" class="col-sm-12 form-label">Name</label>
                <input type="text" class="form-control" :id="id + 'name'" :value="value.name" @input="(v) => {value.name = v.target.value}">
            </div>
            <div class="form-check">
                <input class="form-check-input" type="checkbox" :id="id + 'afk'" :checked="value.afk" @input="(v) => {value.afk = v.target.checked}"">
                <label class="form-check-label" :for="id + 'afk'">AFK</label>
            </div>
        </div>
    </div>
    <span class="badge bg-warning text-dark" v-else>Invalid type "{{ type }}" for "{{ id }}"</span>
    `,
    data: function () {
        return {
            "intents": [
                "GUILDS",
                "GUILD_MEMBERS",
                "GUILD_BANS",
                "GUILD_INTEGRATIONS",
                "GUILD_EMOJIS",
                "GUILD_WEBHOOKS",
                "GUILD_INVITES",
                "GUILD_VOICE_STATES",
                "GUILD_PRESENCES",
                "GUILD_MESSAGES",
                "GUILD_MESSAGE_REACTIONS",
                "GUILD_MESSAGE_TYPING",
                "DIRECT_MESSAGES",
                "DIRECT_MESSAGE_REACTIONS",
                "DIRECT_MESSAGE_TYPING",
            ],
            "selectedIntent": []
        }
    },
    mounted: function () {
        if (this.type == "intent") {
            this.fromIntents(this.value);
        }
    },
    methods: {
        copyFormInputPassword: function () {
            var elem = document.createElement("textarea")
            elem.value = this.$el.lastChild.firstChild.value;
            elem.type = "hidden";
            document.body.append(elem);

            elem.select();
            elem.setSelectionRange(0, 99999);

            document.execCommand("copy");
            elem.parentElement.removeChild(elem);
        },
        updateValue: function (value) {
            this.$emit('input', value)
        },
        calculateIntent() {
            this.intentValue = 0
            this.selectedIntent.forEach(a => { this.intentValue += (1 << a); })
            this.updateValue(Number(this.intentValue))
        },
        fromIntents(val) {
            var _binary = Number(val).toString(2).split("").reverse()
            var _newIntent = [];
            _binary.forEach((value, index) => {
                if (value === "1" && this.selectedIntent.indexOf(value) === -1) {
                    _newIntent.push(index)
                };
            });
            this.selectedIntent = _newIntent;
        },
        updateValue: function (value) {
            this.$emit('input', value)
        }
    },
})

vue = new Vue({
    el: '#app',
    data() {
        return {
            version: "...",
            loading: true,
            error: false,
            daemon: {},
            toast: {
                title: "",
                body: "",
            },
            analytics: {
                chart: {},
                uptime: "...",
                visible: "...",
                events: "...",
                online: "...",
                colour: "bg-success",
            },
            loadingAnalytics: true,

            resttunnel: {
                charts: {
                    hits: {},
                    misses: {},
                    waiting: {},
                    requests: {},
                    callbacks: {},
                    average_response: {}
                },
                numbers: {
                    hits: 0,
                    misses: 0,
                    requests: 0,
                    waiting: 0,
                },
                uptime: "...",
            },
            loadingRestTunnel: true,

            createShardGroupDialogueData: {
                cluster: "",
                autoShard: true,
                shardCount: 1,
                autoIDs: true,
                shardIDs: "",
                startImmediately: true,
            },
            stopShardGroupDialogueData: {
                cluster: "",
                shardgroup: 0,
            },

            createManagerDialogueData: {
                identifier: "",
                persist: true,
                token: "",
                prefix: "",
                client: "",
                channel: "",
            },
            deleteManagerDialogueData: {
                confirm: "",
                cluster: "",
            },
            restartManagerDialogueData: {
                confirm: "",
                cluster: "",
            },

            statusShard: ["Idle", "Waiting", "Connecting", "Connected", "Ready", "Reconnecting", "Closed"],
            colourShard: ["dark", "info", "info", "info", "success", "warn", "secondary"],

            statusGroup: ["Idle", "Starting", "Connecting", "Ready", "Replaced", "Closing", "Closed", "Error"],
            colourGroup: ["dark", "info", "info", "success", "info", "warn", "dark", "danger"],

            colourCluster: ["dark", "info", "info", "success", "warn", "warn", "dark", "danger"],
        }
    },
    filters: {
        pretty: function (value) {
            return JSON.stringify(value, null, 2);
        }
    },
    mounted() {
        this.toastModal = new bootstrap.Toast(document.getElementById("toast"), { delay: 2000 })
        this.fetchConfiguration();
        this.fetchAnalytics();
        this.$nextTick(function () {
            window.setInterval(() => {
                this.fetchAnalytics();
                this.fetchClustersData();
            }, 5000);
        })
    },
    methods: {
        sendRPC(method, params, id) {
            axios
                .post('/api/rpc', {
                    'method': method,
                    'params': params,
                    'id': id,
                })
                .then(result => {
                    err = result.data.error
                    if (err != "") {
                        this.showToast("Error executing " + method, err)
                    } else {
                        this.showToast(method, "Executed successfuly")
                    }
                    return result
                })
                .catch(error => { console.log(error); this.showToast("Exception sending RPC", error); })
        },

        showToast(title, body) {
            this.toast.title = title
            this.toast.body = body
            this.toastModal.show()
        },

        saveClusterSettings(cluster) {
            this.sendRPC("manager:update_settings", cluster.configuration)
        },

        saveDaemonSettings() {
            this.sendRPC("daemon:update_settings", this.daemon.configuration)
        },

        stopShardGroupDialogue(cluster, shardgroup) {
            this.stopShardGroupDialogueModal = new bootstrap.Modal(document.getElementById("stopShardGroupDialogue"), {})

            this.stopShardGroupDialogueData.cluster = cluster
            this.stopShardGroupDialogueData.shardgroup = shardgroup

            this.stopShardGroupDialogueModal.show()
        },
        stopShardGroup() {
            config = Object.assign({}, this.stopShardGroupDialogueData)
            this.sendRPC("shardgroup:stop", config)
            setTimeout(() => this.fetchClustersData(), 1000)

            this.stopShardGroupDialogueModal.hide()
        },
        deleteShardGroup(cluster, shardgroup) {
            config = {
                cluster: cluster,
                shardgroup: shardgroup,
            }
            this.sendRPC("shardgroup:delete", config)
            setTimeout(() => this.fetchClustersData(), 1000)
        },

        refreshGateway(cluster) {
            config = {
                cluster: cluster,
            }
            this.sendRPC("manager:refresh_gateway", config)
            setTimeout(() => this.fetchClustersData(), 1000)
        },

        createManagerDialogue() {
            this.createManagerDialogueModal = new bootstrap.Modal(document.getElementById("createManagerDialogue"), {})
            this.createManagerDialogueData.identifier = ""
            this.createManagerDialogueData.persist = true
            this.createManagerDialogueData.token = ""
            this.createManagerDialogueData.prefix = ""
            this.createManagerDialogueData.client = ""
            this.createManagerDialogueData.channel = ""

            this.createManagerDialogueModal.show()
        },
        createManager() {
            this.sendRPC("manager:create", this.createManagerDialogueData)
            setTimeout(() => this.fetchConfiguration(), 1000)

            this.createManagerDialogueModal.hide()
        },

        deleteManagerDialogue(cluster) {
            this.deleteManagerDialogueModal = new bootstrap.Modal(document.getElementById("deleteManagerDialogue"), {})
            this.deleteManagerDialogueData = {
                "confirm": "",
                "cluster": cluster,
            }

            this.deleteManagerDialogueModal.show()
        },
        deleteManager() {
            this.sendRPC("manager:delete", this.deleteManagerDialogueData)
            setTimeout(() => this.fetchConfiguration(), 1000)

            this.deleteManagerDialogueModal.hide()
        },

        restartManagerDialogue(cluster) {
            this.restartManagerDialogueModal = new bootstrap.Modal(document.getElementById("restartManagerDialogue"), {})
            this.restartManagerDialogueData = {
                "confirm": "",
                "cluster": cluster,
            }

            this.restartManagerDialogueModal.show()
        },
        restartManager() {
            this.sendRPC("manager:restart", this.restartManagerDialogueData)
            setTimeout(() => this.fetchConfiguration(), 1000)

            this.restartManagerDialogueModal.hide()
        },

        createShardGroupDialogue(cluster) {
            this.createShardGroupDialogueModal = new bootstrap.Modal(document.getElementById("createShardGroupDialogue"), {})

            this.createShardGroupDialogueData.cluster = cluster
            this.createShardGroupDialogueData.autoShard = true
            this.createShardGroupDialogueData.shardCount = 1
            this.createShardGroupDialogueData.autoIDs = true
            this.createShardGroupDialogueData.shardIDs = ""
            this.createShardGroupDialogueData.startImmediately = true

            this.createShardGroupDialogueModal.show()
        },
        createShardGroup() {
            config = Object.assign({}, this.createShardGroupDialogueData)
            config.shardCount = Number(config.shardCount)
            this.sendRPC("shardgroup:create", config)
            setTimeout(() => this.fetchClustersData(), 1000)

            this.createShardGroupDialogueModal.hide()
        },

        fetchClustersData() {
            axios
                .get('/api/configuration')
                .then(result => {
                    if (result.data.success == false) { return }
                    if (this.error) { document.location.reload() }

                    this.daemon.rest_tunnel_enabled = result.data.response.rest_tunnel_enabled;
                    clusters = Object.keys(result.data.response.managers)
                    for (mgindex in clusters) {
                        cluster_key = clusters[mgindex]
                        cluster = result.data.response.managers[cluster_key]
                        if (cluster_key in this.daemon.managers) {
                            this.daemon.managers[cluster_key].error = cluster.error
                            this.daemon.managers[cluster_key].shard_groups = cluster.shard_groups
                            this.daemon.managers[cluster_key].gateway = cluster.gateway
                        }
                    }
                })
                .catch(error => { console.log(error); this.showToast("Exception fetching cluster data", error); })
        },
        fetchConfiguration() {
            axios
                .get('/api/configuration')
                .then(result => { this.daemon = result.data.response; this.error = !result.data.success })
                .catch(error => { console.log(error); this.showToast("Exception fetching configuration", error); })
                .finally(() => this.loading = false)
        },
        fetchAnalytics() {
            axios
                .get('/api/analytics')
                .then(result => {
                    if (result.data.success == false) { return }
                    if (this.error) { document.location.reload() }

                    this.analytics = result.data.response;

                    let up = 0
                    let total = 0
                    let guilds = 0
                    this.analytics.colour = "bg-success";

                    clusters = Object.values(this.analytics.clusters)
                    for (mgindex in clusters) {
                        cluster = clusters[mgindex]
                        guilds += cluster.guilds
                        shardgroups = Object.values(cluster.status)
                        for (sgindex in shardgroups) {
                            shardgroupstatus = shardgroups[sgindex]
                            if (2 < shardgroupstatus && shardgroupstatus < 4) {
                                up++
                            }
                            total++
                        }
                    }

                    this.analytics.visible = guilds
                    this.analytics.online = up + "/" + total

                    this.error = this.error | !result.data.success;
                })
                .catch(error => { console.log(error); this.showToast("Exception fetching analytics", error); })
                .finally(() => this.loadingAnalytics = false)
            if (this.daemon.rest_tunnel_enabled) {
                axios
                    .get('/api/resttunnel')
                    .then(result => {
                        if (result.data.success == false) { return }
                        if (this.error) { document.location.reload() }

                        this.resttunnel.charts = result.data.charts;
                        this.resttunnel.uptime = result.data.uptime;
                        this.resttunnel.numbers = result.data.numbers;
                    })
                    .catch(error => { console.log(error); this.showToast("Exception fetching resttunnel", error); })
                    .finally(() => this.loadingRestTunnel = false)
            }
        },
        fromClusters(clusters) {
            _clusters = {}
            Object.entries(clusters).forEach((item) => {
                key = item[0]
                value = item[1]

                shardgroups = Object.values(value.shard_groups)
                if (shardgroups.length == 0) {
                    status = 0
                } else {
                    status = shardgroups.slice(-1)[0].status
                }
                if (value.error != "") {
                    status = 7
                }

                _clusters[key] = {
                    configuration: value.configuration,
                    shardgroups: value.shard_groups,
                    gateway: value.gateway,
                    status: status,
                    error: value.error,
                }
            })
            return _clusters
        },
        calculateAverage(cluster) {
            totalShards = 0;
            totalLatency = 0;

            shardgroups = Object.values(cluster.shardgroups)
            for (sgindex in shardgroups) {
                shardgroup = shardgroups[sgindex]
                if (shardgroup.status < 6) {
                    shards = Object.values(shardgroup.shards)
                    for (shindex in shards) {
                        shard = shards[shindex]
                        latency = (new Date(shard.last_heartbeat_ack) - new Date(shard.last_heartbeat_sent))
                        if (latency > 0) {
                            totalLatency += latency
                            totalShards += 1
                        }
                    }
                }
            }
            return Math.round(totalLatency / totalShards) || '-'
        },
        calculateAverageShardGroup(shardgroup) {
            totalShards = 0;
            totalLatency = 0;

            shards = Object.values(shardgroup.shards)
            for (shindex in shards) {
                shard = shards[shindex]
                latency = (new Date(shard.last_heartbeat_ack) - new Date(shard.last_heartbeat_sent))
                if (latency > 0) {
                    totalLatency += latency
                    totalShards += 1
                }
            }
            return Math.round(totalLatency / totalShards) || '-'
        },

        since(uptime) {
            uptime = new Date(uptime)
            difference = (new Date().getTime() - uptime) / 1000

            output = ""
            if (difference > 86400) {
                var days = Math.trunc(difference / 86400);
                if (days > 0) {
                    output += days + "d"
                }
                var difference = difference % 86400;
            }
            if (difference > 3600) {
                var hours = Math.trunc(difference / 3600);
                if (hours > 0) {
                    output += hours + "h"
                }
                var difference = difference % 3600;
            }
            if (difference > 60) {
                var minutes = Math.trunc(difference / 60);
                if (minutes > 0) {
                    output += minutes + "m"
                }
                var seconds = Math.ceil(difference % 60);
                if (seconds > 0) {
                    output += seconds + "s"
                }
            }
            return output
        }
    }
})