package main

import (
	"github.com/markbates/pkger"
	"github.com/markbates/pkger/pkging/mem"
)

var _ = pkger.Apply(mem.UnmarshalEmbed([]byte(`1f8b08000000000000ffec5add739b38b4ff573a7af6962f93c47e8bd342ecd6ecc6698c61a7d31148368a25c122b04d76fabfdf11fec2894c7b3b776ff390e934e81c1d099d4ffd84f52f207c9e0ad0ff17f084cac70792833ed01e04ce85765fe015d63ee46916a51b6d916a228f3529d8014396a579f1172c12d0075b960719067dc020e1a0033ea431e803d0015f60bec045cba411e1bb492769da22083a600c8b3801fdbfc17bf0b503ee0b4831e817798977c404439172b9a2b47847b82820a518bd8bcae21d5c41426144f13bc2df4525a1e85d0ce304830e705387502ce4bc3ca1ef17a99c7cab60cdc43c4e11e10bed514ede01730a17f2c10ad00124d5485a1644ae9fe3424b8a22031d900ad00122cda5882872c21792511086e5e41f7056cf1c957322f58aaa02cbfe3865598e85d0e61416b8c9583c91aca6790109c7b9468928760cbca95b799515e9a1a1c1ed8c5b22265982f3238d9a9d48c02381e3531299b66df45e3034c20b9c7348358cd63047e2b918a5242b487ce4240c36a8c3f01c72b433dff32e514605c5c70e86ec2321c735a8b8db209a0a88041a2794695f9cd0b66136e867af2c68c34e1b5bef9d525ab6241bd0394448a3a941c18d261d41812fba271cc2615e3539096eceb68fb7039d6126c93c4f73f13c0e132892dd438bf3d8924a1d2c29e30bd24593156765939cb36217ad0716c74591c3183779a9a8dfde646529a54dfaf9901ccf298e0b4a8a13b6207c41f19c924572f25651891852aa95bcb6ec815f6051d0b456217d96765b3693b568fbd022b238360bb16fef6286c92cdc3e3456d28264b0d6bc66fc53a60546594e7821abc536ab4f737bdfacffec4d7460ee57bce3c9ecccf2b4ce4c499739dd978754687883655264dbb5cb87362714efe89de964abe4c576d5bb9616d705714f1dcc048b94d559f7a267679f177c519d96aa38e5ab93a2252a1eef1ec7e9776edad7b40e28398953d468696531372e4ee9ab9a14702ee55698a334d71629857cf13ecd17da46db97ab04c60934f52ca59561e9f60fa4eb878cf19f95db27679b7099aff0bea8b5c8254b346f9738a6dc4eab9f143fd4a216e11f1848061ce242fe675808b83837dd49482fca3a657e2897e5e9a6fa81a0a925198c972d520471d8d29da7e559bf8a4aecea98aa57669e26705ce6588b08227979d69ab56891432ee669ceda84f6c12c27fc19392ee7fbda015fb0284e71c63759d1766043f6360087ec91bbc9d70e18a7482ebbff2f38c15a6309b276c8e767219b9b8e53f463496d91be6729aa074c712e480da88cf78605be7fffde01f3ad0e3562ec6b0c1710c1026a05864c36ded79b567f8b2ce513e10212ba1db14588a7a21d20c813067dbb7bd1eb00266b48bf6bea75f35b5d5cfac0d44dfd0fddf843ef7ed1f5be6ef58d8bf77ad7bcec5976b7fb876ef7759954447c43d216734805ae4b548d68f10af48d8bcbcb4b53ce3aae5f6059c6d555077894f025e81b1d30e429e8db9756d7322ead0e782008f46ddde80057b6e4c0099213e91df01744dfe245fa4d07fdbff54efdef6b075ccba58a0cc7f29df7f261d897578661eb9776077842722e7b1757e6e5a5f9bd03c667c4bb46ef622fbe57f07b07dcfcefc407242f923343f4cbabee7e8861da5df34ab7be77c07dc309039ac64b21ad26dbcbad7fbaba7c9543e142d4667031af9f9f4506735c37ef764d6996af75b048d0f9b3a78a432c3d3f5e34ce16478997e78be301e27852e025a58783c22e2e4e4f0a4dfcbf953e24e296dc62f56dfb248b8fac67235e59ceeee0855cce72f1fcb4762e81bf77406de83ef0d7d9e50db95e0c195d0e1fd3c5f87efd6978734d027f44623635a1efe8d0ef95b26fc8ec7fc29ba1d8ca7b19ba5d923fc935f9c2a6d6f0634663ee517cbfef7704f49da77adc724211a322aef67d930c91e1c5f0665484b34915598355e4523df0bbe4f375ba18de8c9670163e41b757a19b86dc61bcb78a584863362d03f3a15e43e83b4fc89d5611d9c9708f46eef411b9d35aaf903922361f3ecdefd79ff0ba7e4716cab93f1cd64b023e92eb4c904b57d1767d25f697fb353dc2d9c4deea6353540dc2c81ae543b2b7d7260b66e37a2df7334f04bebd0c67a3a77d7fe852736b4383a2dbd12a767b02ce26c961bdd226dc4b236ba437e50ef39bbd3264d32af4edc7f0be5e1b0d669e1ece46e55e2636a722f43d3df4efc89f64c0027ff314de65bdcfd7cf7c5cb5fad88ec9191f3b741533b1b74712b93489ab7a2df7816ff3d7aaaf77d3a22f1fa4f0fea5be0f6e9245aeb30cfdcd1aba3451f979c236f6abd5b92d8ff920438a3c7e70a98e6e3d8266230e6f14794c073462c880befd3a75266d713d5a1debcf51e73bd67b426eaf54f9f78e8fcee8eae891eb94c8a58fcdbed84c5631bfabc706b389f84f746ccddd91112a72f78e4f99ac7fc7fad9f4a997c8baa7a8cd09723749c42759509df3a5a34737ff454d5eb4d5e402b96b454d3656119f54eb0b9a44aad85d7a49c49c25f4a7e5ebd275d9a6eb0addde2974ede9c875ac40b5d7528f4672ad96aa2efdbeb81d7fb96ed3538face14b3de96415b35e89dc2bd59ee305b38108ef07e3d077d8ebf2e9c76abfde808daa701666c16c924566776b63d3a98e36a63adedaf643301b491c542afdba4c8c988fb2c074ca70366ee0a4a09e33f2a77acc7aebc8dc64c875ca73feddca1decf91899360b67231a318f6e6d324d626b426376c0624fa1bfa18135a1b56e6e9844b71ebde1bab8e1a27e07ac7d7fbd69c715a10814fbec846d5631a34b752c0fa4ff52f51efb3b63f9ae359603a6a84f749244b34132fc384ad41879938566a2478c96e1eb8a65a32d9691e5a58a589e05332f85becdcfeeada6b3ae71c2edeb8ce5762ce12931e39d99c8fce4ca7d76391281efe5d0752c68bdd43930a725ba1d35f6aee7fef5e81153bdf42f32a7bfea5fabcdbfa13bd515fefd589fc74caaab7cfbc0a6cba183b288a15796b7adb8a28ccd85eaac97c46662427f2342758daa427f9220b75729f6a0c7d0b7f59839e2acaee6b46cd1d50a7f59d7565cf1a4c450d4d33f93c138b2a6597c2686913f55f8d4a3127b05fefadcf9fd313c9eed5fe62947f417f3b41ab79ded98c7b1324f1d119a4ea5aec3a1087c438587d781ef65213b1fb7b1eb3cfe27be7cfcd8e6cbc7c854e1616f15b9bd2af0272b65fd9d8589ccc3c04ce8abf2676bdd9d2e231586f0274564dabaac590a9cf817f43734e62f6baef417f4c3f37b2a1f2441cb37a75fafb9c3a7d63d9539a5a2e6fa816f3f46d6d40c6743959e1f02df3602b378999f520f9332e88fcff9731d1cfbfe4ffdd9faed8539a50a13def97602dd6916a9bf25be3e1d5bbf994e1a7b7a23665d4744aef3a4dc3fad49a28ad7df871186665bbc46ae932be2751c59e3c59d6f736947f57e52a831c26ff4e59faddf0abd7f54df901e4ca71c7ecc56b1f9a0cacb2fd0752a15e6fb8df5c76e3f9f8efe519d4f23736344fec8388f0d1e16f7cc7984e6547f9ddffac7ebd638f6a9aaee7ad0b7cb70e6ad90fb51e5df19f437cb57b4871aad7b2847eb5011c353939611a3ebd08f553a7e0a6793d784fb8cf66f0d4ea53c9fcd461966bd33df57bc15f67bbaeadbcaefcbd307bd1d272055ddf5431f25ea1c45abc89dd0880c3e478c7278abf0e9ff5fedfd54df89ca312f8ebf5e1e7fba04fb5f98f1dbc5d5b78bab6f1757df2eaebe5d5c7dbbb8fa7671f5ede2eadbc5d5b78babcf2faefe0f000000ffff010000ffff3877d4caeb340000`)))