import { TypographyAtom } from "../components/atoms/TypographyAtom";
import { MahasiswaForm } from "../components/organisms/MahasiswaForm";
import { postMahasiswa } from "../services/mahasiswaService";
import { useNavigate } from "react-router-dom";
import Swal from "sweetalert2";
export function TambahMahasiswaPage() {
    const navigate = useNavigate();

    const handle_cancel = () => { 
        // Kembali ke halaman mahasiswa tanpa menyimpan data
        navigate("/mahasiswa");
    }

// ...existing code...

const handleSubmit = async (data) => {
    try {
        console.log("Mengirim data:", data); // Debug log
        const result = await postMahasiswa(data);
        console.log("Response:", result); // Debug log
        
        Swal.fire({
            icon: "success",
            title: "Berhasil!",
            text: "Data mahasiswa berhasil disimpan.",
            confirmButtonColor: "#22c55e",
        }).then(() => {
            navigate("/mahasiswa");
        });
    } catch (error) {
        console.error("Error saat menyimpan:", error); // Debug log
        
        let errorMessage = "Gagal menyimpan data mahasiswa.";
        
        if (error.response?.data?.message) {
            errorMessage += ` ${error.response.data.message}`;
        } else if (error.response?.data?.error) {
            errorMessage += ` ${error.response.data.error}`;
        }
        
        Swal.fire({
            icon: "error",
            title: "Gagal!",
            text: errorMessage,
            confirmButtonColor: "#ef4444",
        });
    }
};

// ...existing code...

    return (
        <div className="py-6 px-4">
            <TypographyAtom variant="h5" className="mb-4">
                Tambah Data Mahasiswa
            </TypographyAtom>
            <MahasiswaForm onSubmit={handleSubmit} onCancel={ handle_cancel} />
        </div>
    );
}