CREATE TABLE sidangs (
                        id INT AUTO_INCREMENT PRIMARY KEY,
                        mahasiswa_id INT,
                        pembimbing1_id INT,
                        pembimbing2_id INT,
                        judul VARCHAR(255),
                        eprt INT,
                        doc_ta VARCHAR(255),
                        makalah VARCHAR(255),
                        tak INT,
                        status VARCHAR(255),
                        status_pembimbing1 BOOLEAN,
                        status_pembimbing2 BOOLEAN,
                        sks_lulus INT,
                        sks_belum_lulus INT,
                        is_english BOOLEAN,
                        period VARCHAR(255),
                        sk_penguji VARCHAR(255),
                        form_bimbingan1 INT,
                        form_bimbingan2 INT,
                        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
