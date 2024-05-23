USE cool_data;

CREATE TABLE coffee_review
(
    id          BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    coffee_name VARCHAR(50),
    origin      VARCHAR(50),
    review      TEXT,
    created_at  TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = INNODB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8;


INSERT INTO coffee_review (coffee_name, origin, review)
VALUES ('Yirgacheffe Single Origin', 'Etiópia',
        'Assim que eu levei a xícara aos lábios, parecia que eu estava num mercado de especiarias. O aroma floral e cítrico tomou conta e me deixou super curioso. No primeiro gole, senti uma explosão de sabores complexos: limão e bergamota, com um toque sutil de jasmim. A acidez brilhante equilibrada por uma doçura suave, quase como mel, foi incrível. A cada gole, o café ia mostrando mais camadas, deixando um gostinho refrescante na boca. É como uma dança de sabores no paladar, perfeito para começar o dia cheio de energia e inspiração.'),
       ('Blend Especial Fazenda Santa Helena', 'Minas Gerais, Brasil',
        'Desde que o espresso começou a sair, o cheiro rico e achocolatado tomou conta do lugar. O primeiro gole foi como um abraço quentinho - notas de chocolate amargo, nozes torradas e um toque de caramelo. A crema espessa dá uma textura cremosa, deixando tudo ainda mais gostoso. A doçura natural e o corpo encorpado desse blend são bem típicos dos cafés brasileiros. Cada gole traz uma sensação de conforto e familiaridade, perfeito para uma pausa à tarde ou um bate-papo com amigos.'),
       ('Sierra Nevada Organic', 'Colômbia',
        'Quando comecei a prensar o café na Prença Francesa, o cheiro doce e terroso começou a se espalhar pelo ambiente. No primeiro gole, já senti uma profundidade incrível: notas de frutas vermelhas, como cereja e amora, com um toque leve de cacau. A acidez suave e a doçura das frutas se equilibram perfeitamente, e a textura é super aveludada. Esse café me fez sentir super conectado com a natureza, como se eu estivesse saboreando os frutos direto da terra. É uma experiência revigorante e autêntica, perfeita para quem curte um café cheio de personalidade.'),
       ('Café rei do samba', 'Diversas',
        'Preparei esse café na máquina automática e o aroma que surgiu foi ok, nada demais. No primeiro gole, o sabor era simples e direto, sem muita complexidade. Tinha notas de nozes levemente tostadas e um toque sutil de chocolate ao leite. A acidez estava bem equilibrada e o corpo era médio, dando uma sensação satisfatória, mas sem impressionar. Esse café não tem nada de muito especial, mas é uma escolha segura pra quem quer um café consistente e previsível, sem grandes surpresas.'),
       ('Café Top Completo', 'Alto da Compadecida, Brasil',
        'Ao coar esse café, o aroma foi bem fraquinho e nada convidativo. O primeiro gole foi uma decepção - um sabor raso e sem vida, com um amargor queimado que lembrava borra de café velha. A acidez estava desbalanceada e a textura era bem aquosa, deixando uma sensação desagradável na boca. Não havia qualquer traço de doçura ou complexidade para salvar a experiência. Cada gole foi um esforço para terminar a xícara, e fiquei frustrado e arrependido de não ter escolhido um café melhor.');