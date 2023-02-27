-- -------------------------------------------------------------
-- TablePlus 5.3.0(486)
--
-- https://tableplus.com/
--
-- Database: sandbox
-- Generation Time: 2023-02-27 23:29:34.5040
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


CREATE TABLE `codes` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `code` int DEFAULT NULL,
                         `user_id` int DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `expire_at` timestamp NULL DEFAULT NULL,
                         `action` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `data` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `is_used` tinyint(1) NOT NULL DEFAULT '0',
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `comments` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `post_id` int DEFAULT NULL,
                            `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                            `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                            `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2051 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `posts` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `user_id` int DEFAULT NULL,
                         `title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `body` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `email` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `password` varchar(60) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT NULL,
                         `updated_at` timestamp NULL DEFAULT NULL,
                         `role` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `is_active` tinyint(1) DEFAULT '0',
                         `google_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `codes` (`id`, `code`, `user_id`, `created_at`, `expire_at`, `action`, `data`, `is_used`) VALUES
                                                                                                          (1, 620590, 1, '2023-02-11 09:51:13', '2023-02-11 10:51:13', 'registration', '', 1),
                                                                                                          (2, 614029, 2, '2023-02-11 09:54:51', '2023-02-11 10:54:51', 'registration', '', 1);

INSERT INTO `comments` (`id`, `post_id`, `name`, `email`, `body`) VALUES
                                                                      (302, 61, 'quia voluptatem sunt voluptate ut ipsa', 'Lindsey@caitlyn.net', 'fuga aut est delectus earum optio impedit qui excepturi\niusto consequatur deserunt soluta sunt\net autem neque\ndolor ut saepe dolores assumenda ipsa eligendi'),
                                                                      (303, 61, 'excepturi itaque laudantium reiciendis dolorem', 'Gregory.Kutch@shawn.info', 'sit nesciunt id vitae ut itaque sapiente\nneque in at consequuntur perspiciatis dicta consequatur velit\nfacilis iste ut error sed\nin sequi expedita autem'),
                                                                      (304, 61, 'voluptatem quidem animi sit est nemo non omnis molestiae', 'Murphy.Kris@casimer.me', 'minus ab quis nihil quia suscipit vel\nperspiciatis sunt unde\naut ullam quo laudantium deleniti modi\nrerum illo error occaecati vel officiis facere'),
                                                                      (305, 61, 'non cum consequatur at nihil aut fugiat delectus quia', 'Isidro_Kiehn@cristina.org', 'repellendus quae labore sunt ut praesentium fuga reiciendis quis\ncorporis aut quis dolor facere earum\nexercitationem enim sunt nihil asperiores expedita\neius nesciunt a sit sit'),
                                                                      (306, 62, 'omnis nisi libero', 'Kenton_Carter@yolanda.co.uk', 'ab veritatis aspernatur molestiae explicabo ea saepe molestias sequi\nbeatae aut perferendis quaerat aut omnis illo fugiat\nquisquam hic doloribus maiores itaque\nvoluptas amet dolorem blanditiis'),
                                                                      (307, 62, 'id ab commodi est quis culpa', 'Amos_Rohan@mozelle.tv', 'sit tenetur aut eum quasi reiciendis dignissimos non nulla\nrepellendus ut quisquam\nnumquam provident et repellendus eum nihil blanditiis\nbeatae iusto sed eius sit sed doloremque exercitationem nihil'),
                                                                      (308, 62, 'enim ut optio architecto dolores nemo quos', 'Timothy_Heathcote@jose.name', 'officiis ipsa exercitationem impedit dolorem repellat adipisci qui\natque illum sapiente omnis et\nnihil esse et eum facilis aut impedit\nmaxime ullam et dolorem'),
                                                                      (310, 62, 'sed qui atque', 'Toni@joesph.biz', 'quae quis qui ab rerum non hic\nconsequatur earum facilis atque quas dolore fuga ipsam\nnihil velit quis\nrerum sit nam est nulla nihil qui excepturi et'),
                                                                      (311, 63, 'veritatis nulla consequatur sed cumque', 'Brisa@carrie.us', 'officia provident libero explicabo tempora velit eligendi mollitia similique\nrerum sit aut consequatur ullam tenetur qui est vero\nrerum est et explicabo\nsit sunt ea exercitationem molestiae'),
                                                                      (312, 63, 'libero et distinctio repudiandae voluptatem dolores', 'Jasen.Kihn@devon.biz', 'ipsa id eum dolorum et officiis\nsaepe eos voluptatem\nnesciunt quos voluptas temporibus dolores ad rerum\nnon voluptatem aut fugit'),
                                                                      (313, 63, 'quia enim et', 'Efren.Konopelski@madisyn.us', 'corporis quo magnam sunt rerum enim vel\nrepudiandae suscipit corrupti ut ab qui debitis quidem adipisci\ndistinctio voluptatibus vitae molestias incidunt laboriosam quia quidem facilis\nquia architecto libero illum ut dicta'),
                                                                      (314, 63, 'enim voluptatem quam', 'Demetris.Bergnaum@fae.io', 'sunt cupiditate commodi est pariatur incidunt quis\nsuscipit saepe magnam amet enim\nquod quis neque\net modi rerum asperiores consequatur reprehenderit maiores'),
                                                                      (315, 63, 'maxime nulla perspiciatis ad quo quae consequatur quas', 'Luella.Pollich@gloria.org', 'repudiandae dolores nam quas\net incidunt odio dicta eum vero dolor quidem\ndolorem quisquam cumque\nmolestiae neque maxime rerum deserunt nam sequi'),
                                                                      (316, 64, 'totam est minima modi sapiente nobis impedit', 'Sister.Morissette@adelia.io', 'consequatur qui doloribus et rerum\ndebitis cum dolorem voluptate qui fuga\nnecessitatibus quod temporibus non voluptates\naut saepe molestiae'),
                                                                      (317, 64, 'iusto pariatur ea', 'Shyanne@rick.info', 'quam iste aut molestiae nesciunt modi\natque quo laudantium vel tempora quam tenetur neque aut\net ipsum eum nostrum enim laboriosam officia eligendi\nlaboriosam libero ullam sit nulla voluptate in'),
                                                                      (318, 64, 'vitae porro aut ex est cumque', 'Freeman.Dare@ada.name', 'distinctio placeat nisi repellat animi\nsed praesentium voluptatem\nplaceat eos blanditiis deleniti natus eveniet dolorum quia tempore\npariatur illum dolor aspernatur ratione tenetur autem ipsum fugit'),
                                                                      (319, 64, 'et eos praesentium porro voluptatibus quas quidem explicabo est', 'Donnell@orland.org', 'occaecati quia ipsa id fugit sunt velit iure adipisci\nullam inventore quidem dolorem adipisci optio quia et\nquis explicabo omnis ipsa quo ut voluptatem aliquam inventore\nminima aut tempore excepturi similique'),
                                                                      (320, 64, 'fugiat eos commodi consequatur vel qui quasi', 'Robin@gaylord.biz', 'nihil consequatur dolorem voluptatem non molestiae\nodit eum animi\nipsum omnis ut quasi\nvoluptas sed et et qui est aut'),
                                                                      (321, 65, 'rem ducimus ipsam ut est vero distinctio et', 'Danyka_Stark@jedidiah.name', 'ea necessitatibus eum nesciunt corporis\nminus in quisquam iste recusandae\nqui nobis deleniti asperiores non laboriosam sunt molestiae dolore\ndistinctio qui officiis tempora dolorem ea'),
                                                                      (322, 65, 'ipsam et commodi', 'Margarita@casper.io', 'id molestiae doloribus\nomnis atque eius sunt aperiam\ntenetur quia natus nihil sunt veritatis recusandae quia\ncorporis quam omnis veniam voluptas amet quidem illo deleniti'),
                                                                      (323, 65, 'et aut non illo cumque pariatur laboriosam', 'Carlo@cortney.net', 'explicabo dicta quas cum quis rerum dignissimos et\nmagnam sit mollitia est dolor voluptas sed\nipsum et tenetur recusandae\nquod facilis nulla amet deserunt'),
                                                                      (324, 65, 'ut ut architecto vero est ipsam', 'Mina@nikita.tv', 'ipsam eum ea distinctio\nducimus saepe eos quaerat molestiae\ncorporis aut officia qui ut perferendis\nitaque possimus incidunt aut quis'),
                                                                      (325, 65, 'odit sit numquam rerum porro dolorem', 'Violette@naomi.tv', 'qui vero recusandae\nporro esse sint doloribus impedit voluptatem commodi\nasperiores laudantium ut dolores\npraesentium distinctio magnam voluptatum aut'),
                                                                      (326, 66, 'voluptatem laborum incidunt accusamus', 'Lauren.Hodkiewicz@jarret.info', 'perspiciatis vero nulla aut consequatur fuga earum aut\nnemo suscipit totam vitae qui at omnis aut\nincidunt optio dolorem voluptatem vel\nassumenda vero id explicabo deleniti sit corrupti sit'),
                                                                      (327, 66, 'quisquam necessitatibus commodi iure eum', 'Ernestina@piper.biz', 'consequatur ut aut placeat harum\nquia perspiciatis unde doloribus quae non\nut modi ad unde ducimus omnis nobis voluptatem atque\nmagnam reiciendis dolorem et qui explicabo'),
                                                                      (328, 66, 'temporibus ut vero quas', 'Hettie_Morar@wiley.info', 'molestiae minima aut rerum nesciunt\nvitae iusto consequatur architecto assumenda dolorum\nnon doloremque tempora possimus qui mollitia omnis\nvitae odit sed'),
                                                                      (329, 66, 'quasi beatae sapiente voluptates quo temporibus', 'Corbin.Hilll@modesto.biz', 'nulla corrupti delectus est cupiditate explicabo facere\nsapiente quo id quis illo culpa\nut aut sit error magni atque asperiores soluta\naut cumque voluptatem occaecati omnis aliquid'),
                                                                      (330, 66, 'illo ab quae deleniti', 'Kenyatta@renee.io', 'dolores tenetur rerum et aliquam\nculpa officiis ea rem neque modi quaerat deserunt\nmolestias minus nesciunt iusto impedit enim laborum perferendis\nvelit minima itaque voluptatem fugiat'),
                                                                      (331, 67, 'nemo cum est officia maiores sint sunt a', 'Don@cameron.co.uk', 'maxime incidunt velit quam vel fugit nostrum veritatis\net ipsam nisi voluptatem voluptas cumque tempora velit et\net quisquam error\nmaiores fugit qui dolor sit doloribus'),
                                                                      (332, 67, 'dicta vero voluptas hic dolorem', 'Jovan@aaliyah.tv', 'voluptas iste deleniti\nest itaque vel ea incidunt quia voluptates sapiente repellat\naut consectetur vel neque tempora esse similique sed\na qui nobis voluptate hic eligendi doloribus molestiae et'),
                                                                      (333, 67, 'soluta dicta pariatur reiciendis', 'Jeanie.McGlynn@enoch.ca', 'et dolor error doloremque\nodio quo sunt quod\nest ipsam assumenda in veniam illum rerum deleniti expedita\nvoluptate hic nostrum sed dolor et qui'),
                                                                      (334, 67, 'et adipisci laboriosam est modi', 'Garett_Gusikowski@abigale.me', 'et voluptatibus est et aperiam quaerat voluptate eius quo\nnihil voluptas doloribus et ea tempore\nlabore non dolorem\noptio consequatur est id quia magni voluptas enim'),
                                                                      (335, 67, 'voluptatem accusantium beatae veniam voluptatem quo culpa deleniti', 'Doug@alana.co.uk', 'hic et et aspernatur voluptates voluptas ut ut id\nexcepturi eligendi aspernatur nulla dicta ab\nsuscipit quis distinctio nihil\ntemporibus unde quasi expedita et inventore consequatur rerum ab'),
                                                                      (336, 68, 'eveniet eligendi nisi sunt a error blanditiis et ea', 'Yoshiko@viviane.name', 'similique autem voluptatem ab et et\nodio animi repellendus libero voluptas voluptas quia\nlibero facere saepe nobis\nconsequatur et qui non hic ea maxime nihil'),
                                                                      (337, 68, 'beatae esse tenetur aut est', 'Micaela_Bins@mertie.us', 'est ut aut ut omnis distinctio\nillum quisquam pariatur qui aspernatur vitae\ndolor explicabo architecto veritatis ipsa et aut est molestiae\nducimus et sapiente error sed omnis'),
                                                                      (338, 68, 'qui sit quo est ipsam minima neque nobis', 'Loy@gillian.me', 'maiores totam quo atque\nexplicabo perferendis iste facilis odio ab eius consequatur\nsit praesentium ea vitae optio minus\nvoluptate necessitatibus omnis itaque omnis qui'),
                                                                      (339, 68, 'accusantium et sit nihil quibusdam voluptatum provident est qui', 'Tyrel@hunter.net', 'dicta dolorem veniam ipsa harum minus sequi\nomnis quia voluptatem autem\nest optio doloribus repellendus id commodi quas exercitationem eum\net eum dignissimos accusamus est eaque doloremque'),
                                                                      (340, 68, 'rerum et quae tenetur soluta voluptatem tempore laborum enim', 'Otilia.Schuppe@randal.com', 'est aut consequatur error illo ut\nenim nihil fuga\nsuscipit inventore officiis iure earum pariatur temporibus in\naperiam qui quod vel necessitatibus velit eos exercitationem culpa'),
                                                                      (341, 69, 'sunt ut voluptatem cupiditate maxime dolores eligendi', 'April@larissa.co.uk', 'iure ea ea neque est\nesse ab sed hic et ullam sed sequi a\nnon hic tenetur sunt enim ea\nlaudantium ea natus'),
                                                                      (342, 69, 'corporis at consequuntur consequatur', 'Glenna_Waters@retha.me', 'quis beatae qui\nsequi dicta quas et dolor\nnon enim aspernatur excepturi aut rerum asperiores\naliquid animi nulla ea tempore esse'),
                                                                      (343, 69, 'repellat sed consequatur suscipit aliquam', 'Cordelia.Oberbrunner@peyton.com', 'ea alias eos et corrupti\nvoluptatem ab incidunt\nvoluptatibus voluptas ea nesciunt\ntotam corporis dolor recusandae voluptas harum'),
                                                                      (344, 69, 'blanditiis rerum voluptatem quaerat modi saepe ratione assumenda qui', 'Zander@santino.net', 'iusto nihil quae rerum laborum recusandae voluptatem et necessitatibus\nut deserunt cumque qui qui\nnon et et eos adipisci cupiditate dolor sed voluptates\nmaiores commodi eveniet consequuntur'),
                                                                      (345, 69, 'ut deleniti autem ullam quod provident ducimus enim explicabo', 'Camila_Runolfsdottir@tressa.tv', 'omnis et fugit eos sint saepe ipsam unde est\ndolores sit sit assumenda laboriosam\ndolor deleniti voluptatem id nesciunt et\nplaceat dolorem cumque laboriosam sunt non'),
                                                                      (346, 70, 'beatae in fuga assumenda dolorem accusantium blanditiis mollitia', 'Kirstin@tina.info', 'quas non magnam\nquia veritatis assumenda reiciendis\nsimilique dolores est ab\npraesentium fuga ut'),
                                                                      (347, 70, 'tenetur id delectus recusandae voluptates quo aut', 'Anthony.Koepp@savannah.tv', 'consectetur illo corporis sit labore optio quod\nqui occaecati aut sequi quia\nofficiis quia aut odio quo ad\nrerum tenetur aut quasi veniam'),
                                                                      (348, 70, 'molestias natus autem quae sint qui', 'Bradley.Lang@marilyne.tv', 'perferendis dignissimos soluta ut provident sit et\ndelectus ratione ad sapiente qui excepturi error qui quo\nquo illo commodi\nrerum maxime voluptas voluptatem'),
                                                                      (349, 70, 'odio maiores a porro dolorum ut pariatur inventore', 'Loren@aric.biz', 'dicta impedit non\net laborum laudantium qui eaque et beatae suscipit\nsequi magnam rem dolorem non quia vel adipisci\ncorrupti officiis laudantium impedit'),
                                                                      (350, 70, 'eius quia pariatur', 'Arjun@natalie.ca', 'eaque rerum tempore distinctio\nconsequatur fugiat veniam et incidunt ut ut et\nconsequatur blanditiis magnam\ndoloremque voluptate ut architecto facere in dolorem et aut'),
                                                                      (997, 61, 'excepturi itaque laudantium reiciendis dolorem', 'Gregory.Kutch@shawn.info', 'sit nesciunt id vitae ut itaque sapiente\nneque in at consequuntur perspiciatis dicta consequatur velit\nfacilis iste ut error sed\nin sequi expedita autem'),
                                                                      (998, 61, 'excepturi itaque laudantium reiciendis dolorem', 'Gregory.Kutch@shawn.info', 'sit nesciunt id vitae ut itaque sapiente\nneque in at consequuntur perspiciatis dicta consequatur velit\nfacilis iste ut error sed\nin sequi expedita autem'),
                                                                      (999, 61, 'excepturi itaque laudantium reiciendis dolorem', 'Gregory.Kutch@shawn.info', 'sit nesciunt id vitae ut itaque sapiente\nneque in at consequuntur perspiciatis dicta consequatur velit\nfacilis iste ut error sed\nin sequi expedita autem'),
                                                                      (1000, 61, 'excepturi itaque laudantium reiciendis dolorem', 'Gregory.Kutch@shawn.info', 'sit nesciunt id vitae ut itaque sapiente\nneque in at consequuntur perspiciatis dicta consequatur velit\nfacilis iste ut error sed\nin sequi expedita autem'),
                                                                      (1001, 61, 'excepturi itaque laudantium reiciendis dolorem', 'Gregory.Kutch@shawn.info', 'sit nesciunt id vitae ut itaque sapiente\nneque in at consequuntur perspiciatis dicta consequatur velit\nfacilis iste ut error sed\nin sequi expedita autem'),
                                                                      (1002, 61, 'excepturi itaque laudantium reiciendis dolorem', 'Gregory.Kutch@shawn.info', 'sit nesciunt id vitae ut itaque sapiente\nneque in at consequuntur perspiciatis dicta consequatur velit\nfacilis iste ut error sed\nin sequi expedita autem'),
                                                                      (1003, 61, 'excepturi itaque laudantium reiciendis dolorem', 'Gregory.Kutch@shawn.info', 'sit nesciunt id vitae ut itaque sapiente\nneque in at consequuntur perspiciatis dicta consequatur velit\nfacilis iste ut error sed\nin sequi expedita autem'),
                                                                      (1004, 61, 'excepturi itaque laudantium reiciendis dolorem', 'Gregory.Kutch@shawn.info', 'sit nesciunt id vitae ut itaque sapiente\nneque in at consequuntur perspiciatis dicta consequatur velit\nfacilis iste ut error sed\nin sequi expedita autem'),
                                                                      (2050, 61, 'excepturi itaque laudantium reiciendis dolorem', 'Gregory.Kutch@shawn.info', 'sit nesciunt id vitae ut itaque sapiente\nneque in at consequuntur perspiciatis dicta consequatur velit\nfacilis iste ut error sed\nin sequi expedita autem');

INSERT INTO `posts` (`id`, `user_id`, `title`, `body`) VALUES
                                                           (61, 7, 'voluptatem doloribus consectetur est ut ducimus', 'dolore maxime saepe dolor asperiores cupiditate nisi nesciunt\nvitae tempora ducimus vel eos perferendis\nfuga sequi numquam blanditiis sit sed inventore et\nut possimus soluta voluptas nihil aliquid sed earum'),
                                                           (62, 7, 'voluptatem doloribus consectetur est ut ducimus', 'ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit'),
                                                           (63, 7, 'voluptatem doloribus consectetur est ut ducimus', 'ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit'),
                                                           (66, 7, 'voluptatem doloribus consectetur est ut ducimus', 'dolore maxime saepe dolor asperiores cupiditate nisi nesciunt\nvitae tempora ducimus vel eos perferendis\nfuga sequi numquam blanditiis sit sed inventore et\nut possimus soluta voluptas nihil aliquid sed earum'),
                                                           (999, 8, 'voluptatem doloribus consectetur est ut ducimus', 'ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit'),
                                                           (1000, 7, 'voluptatem doloribus consectetur est ut ducimus', 'ab nemo optio odio delectus tenetur corporis similique nobis repellendus rerum omnis facilis vero blanditiis debitis in nesciunt doloribus dicta dolores magnam minus velit'),
                                                           (1001, 7, 'voluptatem doloribus consectetur est ut ducimus', 'dolore maxime saepe dolor asperiores cupiditate nisi nesciunt\nvitae tempora ducimus vel eos perferendis\nfuga sequi numquam blanditiis sit sed inventore et\nut possimus soluta voluptas nihil aliquid sed earum');

INSERT INTO `users` (`id`, `email`, `password`, `created_at`, `updated_at`, `role`, `is_active`, `google_id`) VALUES
                                                                                                                  (1, '778qwertwy2w4@gmail.com', 'Qwerty12#', '2023-02-11 09:51:13', '2023-02-11 09:51:54', 'user', 1, NULL),
                                                                                                                  (2, '777qwertwy2w4@gmail.com', '$2a$12$qWT0Y3uJr/xG3Q2NOQAbFeeMuS091LIy8WINDgMmjsP1w5QeL4isq', '2023-02-11 09:54:51', '2023-02-11 09:55:01', 'user', 1, NULL),
                                                                                                                  (3, '1artkostpromotion@gmail.com', '', '2023-02-13 16:59:04', '2023-02-13 16:59:04', 'user', 1, '110234705669799849673'),
                                                                                                                  (5, 'artkostpromotion@gmail.com', '', '2023-02-13 17:17:09', '2023-02-13 17:17:09', 'user', 1, '110234705669799849672');



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;